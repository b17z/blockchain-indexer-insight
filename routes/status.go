package routes

import (
	"encoding/json"
	"errors"
    "net/http"
)

type StatusInfoResponseInfo struct {
	Version 			int			`json:"version"`
	ProtocolVersion 	int			`json:"protocolversion"`
	Blocks 				int			`json:"blocks"`
	TimeOffset 			int			`json:"timeoffset"`
	Connections			int			`json:"connections"`
	Proxy 				string		`json:"proxy"`
	Difficulty 			float64		`json:"difficulty"`
	Testnet 			bool		`json:"testnet"`
	RelayFee 			float32		`json:"relayfee"`
	Errors 				string		`json:"errors"`
	Network 			string		`json:"network"`
}

type StatusInfoResponse struct {
	Info 				StatusInfoResponseInfo	`json:"info"`
}

type StatusLastBlockHashResponse struct {
	SyncTipHash			string		`json:"syncTipHash"`
	LastBlockHash		string		`json:"lastblockhash"`
}

func Status(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("q")
  
	switch query {
	case "getInfo":
		StatusInfo(w, req);
	case "getLastBlockHash":
		StatusLastBlockHash(w, req);
	default: 
		http.Error(w, errors.New("Invalid query").Error(), http.StatusInternalServerError)
	}
}

func StatusInfo(w http.ResponseWriter, req *http.Request) {

	status := StatusInfoResponse{StatusInfoResponseInfo{150000,70015,509975,0,8,"",123.456,false,0.001,"","livenet"}}

	js, err := json.Marshal(status)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func StatusLastBlockHash(w http.ResponseWriter, req *http.Request) {

	status := StatusLastBlockHashResponse{"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590","0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590"}

	js, err := json.Marshal(status)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}