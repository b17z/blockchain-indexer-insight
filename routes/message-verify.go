package routes

import (
	"encoding/json"
	"net/http"
)

type MessageVerifyResponse struct {
	Result	bool	`json:"result"`
}

func MessageVerify(w http.ResponseWriter, req *http.Request) {
	response := MessageVerifyResponse{false}

	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
