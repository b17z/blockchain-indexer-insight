package routes

import (
	"encoding/json"
	"net/http"
)

type TxSendResponse struct {
	TransactionId	string	`json:"txid"`
}

func TxSend(w http.ResponseWriter, req *http.Request) {
	response := TxSendResponse{"966331ce06b94cd0b38cc1e411c5146d87c0454727aabef22448d47d24c3a5c0"}

	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
