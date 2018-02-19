package routes

import (
	"encoding/json"
    "net/http"
)

type BlocksResponseBlockPoolInfo struct {
	PoolName	string		`json:"poolName"`
	Url 		string		`json:"url"`
}

type BlocksResponseBlock struct {
	Height		int 							`json:"height"`
	Size		int 							`json:"size"`
	Hash		string							`json:"hash"`
	Time		int64							`json:"time"`
	TxLength	int 							`json:"txlength"`
	PoolInfo	BlocksResponseBlockPoolInfo		`json:"poolInfo"`
}

type BlocksResponse struct {
	Blocks 		[]BlocksResponseBlock 	`json:"blocks"`
	Length 		int						`json:"length"`
}

func Blocks(w http.ResponseWriter, req *http.Request) {
	blocks := []BlocksResponseBlock{BlocksResponseBlock{509977,1156259,"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590", 1519073539, 4392, BlocksResponseBlockPoolInfo{}},BlocksResponseBlock{509976,1156259,"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590", 1519073539, 4392, BlocksResponseBlockPoolInfo{}},BlocksResponseBlock{509975,1156259,"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590", 1519073539, 4392, BlocksResponseBlockPoolInfo{}},BlocksResponseBlock{509974,1156259,"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590", 1519073539, 4392, BlocksResponseBlockPoolInfo{}},BlocksResponseBlock{509973,1156259,"0000000000000000001dc1be57707e1d6df48006e7a6c046f473bad75eca4590", 1519073539, 4392, BlocksResponseBlockPoolInfo{}}}

	response := BlocksResponse{blocks, 5}
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
