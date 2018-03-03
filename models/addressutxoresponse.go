package models

type AddressUtxoResponse struct {
	Address						string					`json:"address"`
	TxId						string					`json:"txid"`
	Vout 						int						`json:"vout"`
	Amount		 				float64					`json:"amount"`
	ScriptPubKey 				string					`json:"scriptPubKey"`
	Satoshis 					int64					`json:"satoshis"`
	Confirmations 				int						`json:"confirmations"`
	Timestamp		 			int64					`json:"ts"`
}