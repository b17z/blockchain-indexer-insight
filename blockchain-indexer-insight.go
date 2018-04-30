package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gertjaap/blockchain-indexer-insight/routes"
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
)

func main() {
	logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	server, err := socketio.NewServer(nil)
	if err != nil {
		logging.Info.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		logging.Info.Println("on connection")
		so.On("subscribe", func(msg string) {
			logging.Info.Println("subscription requested:", msg)
			so.Join(msg)
		})
		so.On("disconnection", func() {
			logging.Info.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		logging.Info.Println("error:", err)
	})

	http.Handle("/socket.io/", server)

	r := mux.NewRouter()
	r.HandleFunc("/addr/{address}", routes.Address)
	r.HandleFunc("/addr/{address}/", routes.Address)
	r.HandleFunc("/addr/{address}/utxo", routes.AddressUtxo)
	r.HandleFunc("/addrs/{addresses}/utxo", routes.MultiAddressUtxo)
	r.HandleFunc("/addrs/utxo", routes.MultiAddressUtxo)
	r.HandleFunc("/addrs/{addresses}/txs", routes.MultiAddressTxs)
	r.HandleFunc("/addrs/txs", routes.MultiAddressTxs)
	r.HandleFunc("/addr/{address}/balance", routes.AddressBalance)
	r.HandleFunc("/addr/{address}/totalReceived", routes.AddressTotalReceived)
	r.HandleFunc("/addr/{address}/totalSent", routes.AddressTotalSent)
	r.HandleFunc("/addr/{address}/unconfirmedBalance", routes.AddressUnconfirmedBalance)

	r.HandleFunc("/blocks", routes.Blocks)
	r.HandleFunc("/block/{hash}", routes.Block)
	r.HandleFunc("/rawblock/{hash}", routes.RawBlock)
	r.HandleFunc("/block-index/{height}", routes.BlockIndex)

	r.HandleFunc("/utils/estimatefee", routes.EstimateFee)
	r.HandleFunc("/currency", routes.Currency)
	r.HandleFunc("/sync", routes.Sync)
	r.HandleFunc("/version", routes.Version)
	r.HandleFunc("/status", routes.Status)
	r.HandleFunc("/peer", routes.Peer)
	r.HandleFunc("/txs", routes.Txs)
	r.HandleFunc("/tx/{txid}", routes.Tx)
	r.HandleFunc("/rawtx/{txid}", routes.RawTx)
	r.HandleFunc("/tx/send", routes.TxSend)
	r.HandleFunc("/messages/verify", routes.MessageVerify)

	http.Handle("/", r)
	config := config.GetConfiguration()

	txTicker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		var mempoolTxes []string
		for range txTicker.C {
			url := fmt.Sprintf("%smempool", config.BackendBaseUrl)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				logging.Error.Println("Error reading mempool from ticker (NewRequest): ", err)
				continue
			}

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				logging.Error.Println("Error reading mempool from ticker (Do): ", err)
				continue
			}

			defer resp.Body.Close()

			var mempool []string
			if err := json.NewDecoder(resp.Body).Decode(&mempool); err != nil {
				logging.Error.Println("Error reading mempool from ticker (Decode JSON): ", err)
				continue
			}

			// detect added txes
			for _, txid := range mempool {
				found := false
				for _, existing := range mempoolTxes {
					if txid == existing {
						found = true
						break
					}
				}

				if !found {
					// new TXID
					tx, err := routes.GetTxFromHost(txid)
					if err != nil {
						logging.Error.Println("Error reading mempool from ticker (GetTxFromHost): ", err)
						break
					}

					js, err := json.Marshal(tx)
					if err != nil {
						logging.Error.Println("Error reading mempool from ticker (Marshal): ", err)
						break
					}
					server.BroadcastTo("inv", "tx", js)
					mempoolTxes = append(mempoolTxes, txid)
				}
			}

			mempoolTxes = mempool
		}
	}()

	blockTicker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		lastBlock := ""
		for range blockTicker.C {
			// get last block from BLKIDX
			url := fmt.Sprintf("%sblocks?limit=1", config.BackendBaseUrl)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				logging.Error.Println("Error getting last blockhash from ticker (NewRequest): ", err)
				continue
			}

			client := &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				logging.Error.Println("Error getting last blockhash from ticker (Do): ", err)
				continue
			}

			defer resp.Body.Close()

			var record routes.BlocksResponse

			if err := json.NewDecoder(resp.Body).Decode(&record.Blocks); err != nil {
				logging.Error.Println("Error getting last blockhash from ticker (Decode JSON): ", err)
				continue
			}

			if len(record.Blocks) > 0 {
				if record.Blocks[0].Hash != lastBlock {
					lastBlock = record.Blocks[0].Hash
					server.BroadcastTo("sync", "block", lastBlock)
				}
			}

		}
	}()

	logging.Info.Println("Using backend base URL:", config.BackendBaseUrl)

	listenErr := http.ListenAndServe(":3000", nil)
	if listenErr != nil {
		logging.Error.Println("ListenAndServe: ", listenErr)
	}
}
