package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gorilla/mux"
)

type BlockResponsePoolInfo struct {
	PoolName string `json:"poolName"`
	Url      string `json:"url"`
}

type BlockResponse struct {
	Hash              string                `json:"hash"`
	Size              int                   `json:"size"`
	Height            int                   `json:"height"`
	Version           int64                 `json:"version"`
	MerkleRoot        string                `json:"merkleroot"`
	Tx                []string              `json:"tx"`
	Time              int64                 `json:"time"`
	Nonce             int64                 `json:"nonce"`
	Bits              string                `json:"bits"`
	Difficulty        float64               `json:"difficulty"`
	ChainWork         string                `json:"chainwork"`
	Confirmations     int                   `json:"confirmations"`
	PreviousBlockHash string                `json:"previousblockhash"`
	Reward            float32               `json:"reward"`
	IsMainChain       bool                  `json:"isMainChain"`
	PoolInfo          BlockResponsePoolInfo `json:"poolInfo"`
}

func Block(w http.ResponseWriter, req *http.Request) {

	config := config.GetConfiguration()
	vars := mux.Vars(req)
	url := fmt.Sprintf("%sblock/%s", config.BackendBaseUrl, vars["hash"])
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Error.Println("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record BlockResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logging.Error.Println("Json decode failed: ", err)
		return
	}

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
