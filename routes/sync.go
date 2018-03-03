package routes

import (
	"encoding/json"
	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
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
	
	config := config.GetConfiguration()
	url := config.BackendBaseUrl + "sync"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logging.Error.Println("NewRequest: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		logging.Error.Println("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record SyncResponse

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logging.Error.Println("Json decode failed: ", err)
		return
	}

	js, err := json.Marshal(record)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
