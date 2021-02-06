package websocket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Client contains the info we have about a connection
type Client struct {
	ID   int
	Conn *websocket.Conn
	Hub  *Hub
}

// Message contains the info we get from the client that needs to be sent
type Message struct {
	Type     string `json:"type"`
	ClientID int    `json:"clientId"`
	Body     string `json:"body"`
}

	// Read will listen in for messages coming from the Clients connection
func (c *Client) Read() {
	defer func() {
		c.Hub.Disconnect <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var response Message
		json.Unmarshal(p, &response)

		message := Message{response.Type, response.ClientID, response.Body}
		c.Hub.Broadcast <- message
	}
}