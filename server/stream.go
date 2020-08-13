package server

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

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

// newHub creates a Hub and returns a pointer to it
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
			// set incoming client connection
			h.connections[client] = true
		case client := <-h.unsubscribe:
			if _, ok := h.connections[client]; ok {
				delete(h.connections, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.connections {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.connections, client)
				}
			}
		}
	}
}

// Connection created for each new connection to the hub
type Connection struct {
	hub *Hub

	// The websocket connection.
	socketConnection *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Connection) readPump() {
	defer func() {
		c.hub.unsubscribe <- c
		c.socketConnection.Close()
	}()
	c.socketConnection.SetReadLimit(maxMessageSize)
	c.socketConnection.SetReadDeadline(time.Now().Add(pongWait))
	c.socketConnection.SetPongHandler(func(string) error { c.socketConnection.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.socketConnection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.socketConnection.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.socketConnection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.socketConnection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.socketConnection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.socketConnection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.socketConnection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
