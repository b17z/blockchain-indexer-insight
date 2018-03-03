package routes

import (
	"encoding/json"
    "net/http"
)

type BlockIndexResponse struct {
	BlockHash	string		`json:"blockHash"`
}

func BlockIndex(w http.ResponseWriter, req *http.Request) {
	response := BlockIndexResponse{"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b"}
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
