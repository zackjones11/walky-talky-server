package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client contains the info we have about a connection
type Client struct {
	ID   int
	Conn *websocket.Conn
}

// Message contains the info we get from the client that needs to be sent
type Message struct {
	Type     string `json:"type"`
	ClientID int    `json:"clientId"`
	Body     string `json:"body"`
}