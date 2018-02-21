package main

import (
	
    "github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gertjaap/blockchain-indexer-insight/routes"
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"net/http"
    "os"
)



func main() {
    logging.Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	server, err := socketio.NewServer(nil)
	if err != nil {
		logging.Info.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		logging.Info.Println("on connection")
		so.Join("chat")
		so.On("chat message", func(msg string) {
			logging.Info.Println("emit:", so.Emit("chat message", msg))
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func() {
			logging.Info.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		logging.Info.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	
	r := mux.NewRouter();

	r.HandleFunc("/sync", routes.Sync)
	r.HandleFunc("/version", routes.Version)
	r.HandleFunc("/status", routes.Status)
	r.HandleFunc("/peer", routes.Peer)
	r.HandleFunc("/blocks", routes.Blocks)
	r.HandleFunc("/block/{hash}", routes.Block)
	r.HandleFunc("/txs", routes.Txs)
	
	http.Handle("/", r);
	
	listenErr := http.ListenAndServe(":3000", nil)
    if listenErr != nil {
        logging.Error.Println("ListenAndServe: ", listenErr)
    }
}