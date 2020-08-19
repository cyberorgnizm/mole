package server

import (
	"github.com/gorilla/websocket"
)

// Peer created for each new connection to the hub
type Peer struct {
	hub *Hub

	// The websocket connection.
	socketConnection *websocket.Conn

	// Buffered channel of outbound messages.
	sink chan []byte
}
