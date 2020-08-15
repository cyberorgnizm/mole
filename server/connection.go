package server

import (
	"github.com/gorilla/websocket"
)

// Connection created for each new connection to the hub
type Connection struct {
	hub *Hub

	// The websocket connection.
	socketConnection *websocket.Conn

	// Buffered channel of outbound messages.
	sink chan []byte
}

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// subscribed connections
	connections map[*Connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Subscribe requests from the connections
	subscribe chan *Connection

	// Unsubscribe requests from the connections
	unsubscribe chan *Connection
}

// Handles multiple concurrent connection
func newHub() *Hub {
	return &Hub{
		connections: make(map[*Connection]bool),
		broadcast:   make(chan []byte),
		subscribe:   make(chan *Connection),
		unsubscribe: make(chan *Connection),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.subscribe:
			// set incoming client connection status in pool
			h.connections[client] = true
		case client := <-h.unsubscribe:
			// removes client connection status from pool
			if _, ok := h.connections[client]; ok {
				delete(h.connections, client)
				close(client.sink)
			}
		case message := <-h.broadcast:
			// broadcast message to all clients in poll
			for client := range h.connections {
				select {
				case client.sink <- message:
				default:
					close(client.sink)
					delete(h.connections, client)
				}
			}
		}
	}
}
