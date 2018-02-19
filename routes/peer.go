package routes

import (
	"encoding/json"
    "net/http"
)

type PeerResponse struct {
	Connected 		bool 	`json:"connected"`
	Host 			string	`json:"host"`
	Port			int		`json:"port"`
}


func Peer(w http.ResponseWriter, req *http.Request) {
	peer := PeerResponse{true,"127.0.0.1",0}

	js, err := json.Marshal(peer)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
