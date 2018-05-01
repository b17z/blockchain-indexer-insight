# blockchain-indexer-insight
Insight compatible API wrapper for Blockchain Indexer

## Status
This thing is under development, and far from ready. Feel free to pitch in.

### Websocket implementation

| Websocket   | Dummy | Working |
|-------------|:-----:|:-------:|
| New TX      | yes   | yes     |
| New block   | yes   | yes     |

### REST API implementation

| API         | Dummy | Working |
|-------------|:-----:|:-------:|
| ```/addr/{address}``` | yes | no |    
| ```/addr/{address}/``` | yes | no |   
| ```/addr/{address}/utxo``` | yes | no |    
| ```/addrs/{addresses}/utxo``` | yes | no |   
| ```/addrs/utxo``` | yes | no |   
| ```/addrs/{addresses}/txs``` | yes | no |  
| ```/addrs/txs``` | yes | no |    
| ```/addr/{address}/balance``` | yes | no |   
| ```/addr/{address}/totalReceived``` | yes | no | 
| ```/addr/{address}/totalSent``` | yes | no |    
| ```/addr/{address}/unconfirmedBalance``` | yes | no |   
| ```/blocks``` | yes | yes |
| ```/block/{hash}``` | yes | yes |    
| ```/rawblock/{hash}``` | yes | no |    
| ```/block-index/{height}``` | yes | no |  
| ```/utils/estimatefee``` | yes | no | 
| ```/currency``` | yes | no | 
| ```/sync``` | yes | yes |   
| ```/version``` | yes | no |
| ```/status``` | yes | no |
| ```/peer``` | yes | no |
| ```/txs?block=&pageNum=``` | yes | yes |
| ```/txs?address=``` | no | no |
| ```/tx/{txid}``` | yes | yes |
| ```/rawtx/{txid}``` | yes | no 
| ```/tx/send``` | yes | no 
| ```/messages/verify``` | yes | no | 


