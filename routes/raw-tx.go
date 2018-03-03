package routes

import (
	"encoding/json"
	"net/http"
)

type RawTxResponse struct {
	RawTx	string		`json:"rawtx"`
}

func RawTx(w http.ResponseWriter, req *http.Request) {
	response := RawTxResponse{"01000000011276a74c111f7bfc09ecc5c6cdb74939c7b879dff157bd7252d7f7403cd3c647010000006a473044022065724574ca1e0a9f4b54f9583045bb61e4ceee4e4ccd463da7fb9e8c95441abe02202d88e2a7a1908e0f1dd4389a9b848e112f29b36e13437373129b0eb974eaea14012103da5bac7b36d5aa38f531c6b9601e21bb598a4b6716ebed38b009a55dabde9440feffffff030000000000000000166a146f6d6e69000000000000001f000000625225bd4022020000000000001976a914d2a43709d8d8f87678748daa74bec82614c473da88aca0437500000000001976a914a25dec4d0011064ef106a983c39c7a540699f22088ac51cf0700"}
	
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
