package main

import (
	
    "github.com/gertjaap/blockchain-indexer-insight/logging"
	"github.com/gertjaap/blockchain-indexer-insight/routes"
	"github.com/googollee/go-socket.io"
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
	
	http.HandleFunc("/sync", routes.Sync)
	http.HandleFunc("/version", routes.Version)
	http.HandleFunc("/status", routes.Status)
	http.HandleFunc("/peer", routes.Peer)
	http.HandleFunc("/blocks", routes.Blocks)
	
	
	listenErr := http.ListenAndServe(":3000", nil)
    if listenErr != nil {
        logging.Error.Println("ListenAndServe: ", listenErr)
    }
}