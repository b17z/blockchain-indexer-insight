package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gertjaap/blockchain-indexer-insight/models"
)

func Txs(w http.ResponseWriter, req *http.Request) {

	block := req.URL.Query().Get("block")
	if block != "" {
		BlockTxs(w, req)
		return
	}

	http.Error(w, "Unexpected query string to /txs : "+req.URL.RawQuery, http.StatusInternalServerError)
}

func BlockTxs(w http.ResponseWriter, req *http.Request) {

	hash := req.URL.Query().Get("block")
	pageNumString := req.URL.Query().Get("block")
	if pageNumString == "" {
		pageNumString = "0"
	}

	pageNum, _ := strconv.ParseUint(pageNumString, 10, 64)
	config := config.GetConfiguration()

	url := fmt.Sprintf("%sblocktxs/%s/%d", config.BackendBaseUrl, hash, pageNum)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Println("Do: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	record := models.TxsResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, tx := range record.Txs {
		tx.ValueIn = float64(0)
		tx.ValueOut = float64(0)

		for _, txi := range tx.Vin {
			txi.Value = float64(txi.ValueSatoshi) / float64(100000000)
			tx.ValueIn += txi.Value
		}

		for _, txo := range tx.Vout {
			txo.Value = float64(txo.ValueSatoshi) / float64(100000000)
			tx.ValueOut += txo.Value
		}

		tx.Fees = tx.ValueIn - tx.ValueOut
	}

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
