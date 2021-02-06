package main

import (
	"fmt"
	"net/http"

	"github.com/zackjones11/walky-talky-server/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	_, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
}


func setupRoutes() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":3000", nil)

}