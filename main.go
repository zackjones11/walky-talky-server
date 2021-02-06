package main

import (
	"fmt"
	"net/http"

	"github.com/zackjones11/walky-talky-server/pkg/websocket"
)

func serveWs(hub *websocket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(err)
	}

	client := &websocket.Client{
		ID:   len(hub.Clients),
		Conn: conn,
		Hub:  hub,
	}

	hub.NewClient <- client

	client.Read()
}


func setupRoutes() {
	hub := websocket.NewHub()
	go hub.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}

func main() {
	setupRoutes()
	http.ListenAndServe(":3000", nil)

}