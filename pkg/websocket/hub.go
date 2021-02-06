package websocket

import "fmt"

// Hub contains all the possible channels
type Hub struct {
	NewClient  chan *Client
	Disconnect chan *Client
	Broadcast  chan Message
	Clients    map[*Client]bool
}

// NewHub lets us use our channels
func NewHub() *Hub {
	return &Hub{
		NewClient:  make(chan *Client),
		Disconnect: make(chan *Client),
		Broadcast:  make(chan Message),
		Clients:    make(map[*Client]bool),
	}
}

// Start is the single entry point to call our channels
func (hub *Hub) Start() {
	for {
		select {
		case client := <-hub.NewClient:
			hub.Clients[client] = true

			for c := range hub.Clients {
				c.Conn.WriteJSON(Message{Type: "NEW_USER", ClientID: c.ID})
			}
			break

		case client := <-hub.Disconnect:
			delete(hub.Clients, client)
			for c := range hub.Clients {
				c.Conn.WriteJSON(Message{Type: "DISCONNECTION", ClientID: c.ID})
			}
			break

		case message := <-hub.Broadcast:
			for c := range hub.Clients {
				if message.ClientID == c.ID {
					continue
				}

				err := c.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}