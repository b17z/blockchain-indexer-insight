package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gertjaap/blockchain-indexer-insight/models"
	"github.com/gorilla/mux"
)

func Tx(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	response, err := GetTxFromHost(vars["txid"])

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetTxFromHost(txid string) (models.Transaction, error) {
	config := config.GetConfiguration()
	response := models.Transaction{}
	response.TxId = txid

	url := fmt.Sprintf("%sgetTransacion/%s", config.BackendBaseUrl, txid)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Error.Println("NewRequest: ", err)
		return response, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Println("Do: ", err)
		return response, err
	}

	defer resp.Body.Close()

	var record models.Transaction

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logging.Error.Println("Json decode failed: ", err)
		return response, err
	}

	record.ValueIn = float64(0)
	record.ValueOut = float64(0)

	for _, txi := range record.Vin {
		record.ValueIn += txi.Value
	}

	for _, txo := range record.Vout {
		record.ValueOut += txo.Value
	}

	record.Fees = record.ValueIn - record.ValueOut

	return response, nil
}
