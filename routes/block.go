package routes

import (
	"encoding/json"
    "net/http"
)

type BlockResponsePoolInfo struct {
	PoolName	string		`json:"poolName"`
	Url 		string		`json:"url"`
}

type BlockResponse struct {
	Hash				string					`json:"hash"`
	Size				int 					`json:"size"`
	Height				int 					`json:"height"`
	Version				int64 					`json:"version"`
	MerkleRoot			string					`json:"merkleroot"`
	Tx					[]string				`json:"tx"`
	Time				int64 					`json:"time"`
	Nonce				int64 					`json:"nonce"`
	Bits				string					`json:"bits"`
	Difficulty			float64					`json:"difficulty"`
	ChainWork			string					`json:"chainwork"`
	Confirmations		int						`json:"confirmations"`
	PreviousBlockHash	string					`json:"previousblockhash"`
	Reward				float32					`json:"reward"`
	IsMainChain			bool					`json:"isMainChain"`
	PoolInfo			BlockResponsePoolInfo	`json:"poolInfo"`
}

func Block(w http.ResponseWriter, req *http.Request) {
	txes := []string{"abc123","abc124"}
	response := BlockResponse{"00000000000000000044a01dc989bf8c097c3270a42eaf77200649ed9830ed8b",319152,510292,536870912,"d212b40161d29d67a5ae938defa19a63806288ad770178bc46e4334a424afc2e",txes,1519243687,954079461,"175d97dc",3007383866429.732,"000000000000000000000000000000000000000001227b33cb93c03402773cfc",1,"0000000000000000004c12e60331a55e7a8a306e4f0231ef3d8d7a2290938db8",12.5,true,BlockResponsePoolInfo{}}
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
