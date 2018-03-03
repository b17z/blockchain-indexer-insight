package routes

import (
	"encoding/json"
    "net/http"
)

type EstimateFeeResponse struct {
	Two 		float64 						`json:"2"`
}


func EstimateFee(w http.ResponseWriter, req *http.Request) {
	response := EstimateFeeResponse{0.000001}

	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
