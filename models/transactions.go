package models

type ScriptSig struct {
	Hex 				string					`json:"hex"`
	Asm  				string 					`json:"asm"`
}

type TransactionVin struct {	
	CoinBase 			string 						`json:"coinbase"`
	Sequence  			int64						`json:"sequence"`
	Number				int 						`json:"n"`
	TxId 				string 						`json:"txid,omitempty"`
	Vout 				int 						`json:"vout,omitempty"`
	ScriptSig 			ScriptSig					`json:"scriptSig,omitempty"`
	Address 			string 						`json:"addr,omitempty"`
	ValueSatoshi		int64						`json:"valueSat,omitempty"`
	Value 				float64 					`json:"value,omitempty"`
	DoubleSpentTxID 	string 						`json:"doubleSpentTxID,omitempty"`
}

type ScriptPubKey struct {
	Hex 				string					`json:"hex"`
	Asm  				string 					`json:"asm"`
	Addresses			[]string 				`json:"addresses"`
	Type 				string 					`json:"type"`
}

type TransactionVout struct {
	Value 				float64							`json:"value"`
	Number 				int 							`json:"n"`
	ScriptPubKey 		ScriptPubKey					`json:"scriptPubKey"`
	SpentTxId 			string							`json:"spentTxId,omitempty"`
	SpentIndex 			int 							`json:"spentIndex,omitempty"`
	SpentHeight 		int  							`json:"spentHeight,omitempty"`
}

type Transaction struct {
	TxId 				string					`json:"txid"`
	Version				int 					`json:"version"`
	LockTime 			int 					`json:"locktime"`
	Vin 				[]TransactionVin 		`json:"vin"`
	Vout 				[]TransactionVout	 	`json:"vout"`
	BlockHash 			string 					`json:"blockhash"`
	BlockHeight 		int 					`json:"blockheight"`
	Confirmations 		int 					`json:"confirmations"`
	Time 				int64 					`json:"time"`
	BlockTime 			int64 					`json:"blocktime"`
	IsCoinBase 			bool 					`json:"isCoinBase"`
	ValueOut 			float64 				`json:"valueOut"`
	Size 				int 					`json:"size"`
	ValueIn 			float64 				`json:"valueIn"`
	Fees	 			float64 				`json:"fees"`	
}

type TxsResponse struct {
	PagesTotal			int 					`json:"pagesTotal"`
	Txs					[]Transaction			`json:"txs"`
}
