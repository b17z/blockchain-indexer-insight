package routes

import (
	"encoding/json"
	"github.com/gertjaap/blockchain-indexer-insight/models"
	"net/http"
)

func Tx(w http.ResponseWriter, req *http.Request) {
	jsonString := `{"txid":"922660cc3296c96ad399f051f29ac3ceb3286acc299382211549246214755e31","version":1,"locktime":0,"vin":[{"txid":"cb632dd5d6d8ba29afc4521fa5645fa06e1425f8fcb54bdce63ad09eeeae4b75","vout":2,"sequence":4294967294,"n":0,"scriptSig":{"hex":"473044022074fd21c6c67b92d8ecfd2a5b3be3900859d558c528ebbc1043f010160b94ee0302200c53a6aac0587b94d873e462a980678049f1d982c381a7e432e52038c28eae3e0121033a41a60347fc2c6567971fff3715aaabda8dcd8e0b9b2ed6e2357b3d6f57ddac","asm":"3044022074fd21c6c67b92d8ecfd2a5b3be3900859d558c528ebbc1043f010160b94ee0302200c53a6aac0587b94d873e462a980678049f1d982c381a7e432e52038c28eae3e01 033a41a60347fc2c6567971fff3715aaabda8dcd8e0b9b2ed6e2357b3d6f57ddac"},"addr":"Vt3M7yWdF7gZkieFyYxM1GxqVmNPQFcsGZ","valueSat":47352592014,"value":473.52592014,"doubleSpentTxID":null,"isConfirmed":null,"confirmations":null,"unconfirmedInput":null}],"valueIn":473.52592014,"fees":0.01000014,"vout":[{"value":230.51592000,"n":0,"scriptPubKey":{"hex":"76a9143243b597befcd0f041228568dddcf752637ad2f888ac","asm":"OP_DUP OP_HASH160 3243b597befcd0f041228568dddcf752637ad2f8 OP_EQUALVERIFY OP_CHECKSIG","addresses":["Veaboq692WBy9RmjX6TQc4rSUpMPqUZLAi"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null},{"value":243.00000000,"n":1,"scriptPubKey":{"hex":"76a91481a87069690287e9aee4253b62edd18d347b974288ac","asm":"OP_DUP OP_HASH160 81a87069690287e9aee4253b62edd18d347b9742 OP_EQUALVERIFY OP_CHECKSIG","addresses":["VmpProziEEdDizqtL9FyAgEyHCoJB6Nb8w"],"type":"pubkeyhash"},"spentTxId":null,"spentIndex":null,"spentHeight":null}],"blockhash":"416caef1dd49b13b3c7e5cbdfbf63a7d407aebc44214730a7f7ffc99fb21c445","blockheight":887003,"confirmations":1,"time":1520095540,"blocktime":1520095540,"valueOut":473.51592,"size":225}`
	response := models.Transaction{}
	json.Unmarshal([]byte(jsonString), &response);
	
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
