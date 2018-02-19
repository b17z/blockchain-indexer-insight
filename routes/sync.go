package routes

import (
	"encoding/json"
    "net/http"
)

type SyncResponse struct {
	BlockChainHeight 	int		`json:"blockChainHeight"`
	Error 				string	`json:"error"`
	Height 				int		`json:"height"`
	Status 				string	`json:"status"`
	SyncPercentage 		float32	`json:"syncPercentage"`
}

func Sync(w http.ResponseWriter, req *http.Request) {
	sync := SyncResponse{1000, "", 1000, "completed", 100.0}

	js, err := json.Marshal(sync)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
