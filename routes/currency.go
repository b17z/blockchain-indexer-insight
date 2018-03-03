package routes

import (
	"encoding/json"
    "net/http"
)

type CurrencyResponseData struct {
	Bitstamp	float64		`json:"bitstamp"`
}

type CurrencyResponse struct {
	Status 		int 						`json:"status"`
	Data 		CurrencyResponseData		`json:"data"`
}


func Currency(w http.ResponseWriter, req *http.Request) {
	response := CurrencyResponse{200,CurrencyResponseData{4.00}}

	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
