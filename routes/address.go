package routes

import (
	"encoding/json"
	"net/http"
)

type AddressResponse struct {
	AddressString				string					`json:"addrStr"`
	Balance						float64					`json:"balance"`
	BalanceSat 					int64					`json:"balanceSat"`
	TotalReceived 				float64					`json:"totalReceived"`
	TotalReceivedSat 			int64					`json:"totalReceivedSat"`
	TotalSent 					float64					`json:"totalSent"`
	TotalSentSat 				int64					`json:"totalSentSat"`
	UnconfirmedBalance 			float64					`json:"unconfirmedBalance"`
	UnconfirmedBalanceSat 		int64					`json:"unconfirmedBalanceSat"`
	UnconfirmedTxApperances     int64                   `json:"unconfirmedTxApperances"`	
	TxApperances     			int64                   `json:"txApperances"`
}

func Address(w http.ResponseWriter, req *http.Request) {
	jsonString := `{"addrStr":"1EXqTWBgQquUjQK3prugZaDHUNvEipQfou","balance":0.04238152,"balanceSat":4238152,"totalReceived":0.2538252,"totalReceivedSat":25382520,"totalSent":0.21144368,"totalSentSat":21144368,"unconfirmedBalance":0,"unconfirmedBalanceSat":0,"unconfirmedTxApperances":0,"txApperances":21}`
	response := AddressResponse{}
	json.Unmarshal([]byte(jsonString), &response);
	
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
