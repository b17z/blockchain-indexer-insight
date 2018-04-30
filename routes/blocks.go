package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gertjaap/blockchain-indexer-insight/config"
	"github.com/gertjaap/blockchain-indexer-insight/logging"
)

type BlocksResponseBlockPoolInfo struct {
	PoolName string `json:"poolName"`
	Url      string `json:"url"`
}

type BlocksResponseBlock struct {
	Height   int                         `json:"height"`
	Size     int                         `json:"size"`
	Hash     string                      `json:"hash"`
	Time     int64                       `json:"time"`
	TxLength int                         `json:"txlength"`
	PoolInfo BlocksResponseBlockPoolInfo `json:"poolInfo"`
}

type BlocksResponse struct {
	Blocks []BlocksResponseBlock `json:"blocks"`
	Length int                   `json:"length"`
}

func Blocks(w http.ResponseWriter, req *http.Request) {

	// If "limit" is specified, return the last X blocks
	limitString := req.URL.Query().Get("limit")
	if limitString != "" {
		limit, err := strconv.ParseInt(limitString, 10, 32)
		if err == nil {
			BlocksByLimit(w, int32(limit))
			return
		}
	}

	// If a date is specified, return blocks of that day
	dateString := req.URL.Query().Get("blockDate")
	if dateString != "" {
		t, err := time.Parse("2006-01-02", dateString)
		if err == nil {
			start := t.Unix()
			end := start + (24 * 60 * 60) - 1
			BlocksByDate(w, start, end)
			return
		}
	}

	// Default: today's blocks
	start := time.Now().Truncate(time.Hour * 24).Unix()
	end := start + (24 * 60 * 60) - 1
	BlocksByDate(w, start, end)
}

func BlocksByLimit(w http.ResponseWriter, limit int32) {

	config := config.GetConfiguration()
	url := fmt.Sprintf("%sblocks?limit=%d", config.BackendBaseUrl, limit)
	RequestAndReturn(url, w)
}

func BlocksByDate(w http.ResponseWriter, start, end int64) {
	config := config.GetConfiguration()
	url := fmt.Sprintf("%sblocksbydate?start=%d&end=%d", config.BackendBaseUrl, start, end)
	RequestAndReturn(url, w)
}

func RequestAndReturn(url string, w http.ResponseWriter) {
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

	var record BlocksResponse

	if err := json.NewDecoder(resp.Body).Decode(&record.Blocks); err != nil {
		logging.Error.Println("Json decode failed: ", err)
		return
	}
	record.Length = len(record.Blocks)

	js, err := json.Marshal(record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
