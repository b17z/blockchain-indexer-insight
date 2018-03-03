package routes

import (
	"encoding/json"
	"github.com/gertjaap/blockchain-indexer-insight/models"
	"net/http"
)

func AddressUtxo(w http.ResponseWriter, req *http.Request) {
	jsonString := `[{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"c02de86395da68c0d1bcb60978d437aa3667a623aa71c87b6dfc53c1485cf5ad","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091999},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"470ec41785b2b9f114cbbb6d118b136aac7c2a3242d7b1fdf5eabbfd63149ddc","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091992},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"d1e7339f424bb0a858b68390fc3168411e051e8be9bd41c41450a564e4f481bd","vout":1,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091987},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"2db0f59ee3cc5d586b7d69c9b4b68eea5aa87cf9700207b03929e56f6b38155a","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091987},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"c12c6bef8658e6c3f8b8e5f04aac864af6f63117d6b604e38ceb784b436cd0f8","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091974},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"f4b089432043dff04fb8fadc2b0bac4cdb5f2aa21ebca0e03c6fa6bfbd01bbd0","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.00000546,"satoshis":546,"confirmations":0,"ts":1520091969},{"address":"1FoWyxwPXuj4C6abqwhjDWdz6D4PZgYRjA","txid":"a550ea8b0fd820342cdf9932e7c0a05dd8671aa3b0425c0d40864d5005e5461f","vout":2,"scriptPubKey":"76a914a25dec4d0011064ef106a983c39c7a540699f22088ac","amount":0.07685024,"satoshis":7685024,"height":511826,"confirmations":1}]`
	response := []models.AddressUtxoResponse{}
	json.Unmarshal([]byte(jsonString), &response);
	
	js, err := json.Marshal(response)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
