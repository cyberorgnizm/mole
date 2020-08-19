package server

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Peer) readPump() {
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
func (c *Peer) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.socketConnection.Close()
	}()
	for {
		select {
		case message, ok := <-c.sink:
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
			n := len(c.sink)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.sink)
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
