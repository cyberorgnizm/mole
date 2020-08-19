package server

import "log"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// subscribed connections
	connections map[*Peer]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Subscribe requests from the connections
	subscribe chan *Peer

	// Unsubscribe requests from the connections
	unsubscribe chan *Peer
}

// Handles multiple concurrent connection
func newHub() *Hub {
	log.Println("[Socket] Creating socket hub...")
	defer log.Println("[Socket] Hub successfully created")
	return &Hub{
		connections: make(map[*Peer]bool),
		broadcast:   make(chan []byte),
		subscribe:   make(chan *Peer),
		unsubscribe: make(chan *Peer),
	}
}

func (h *Hub) run() {
	for {
		log.Println("[Socket] listening on Hub socket channel...")
		select {
		case client := <-h.subscribe:
			// set incoming client connection status in pool
			h.connections[client] = true
			log.Println("[Socket] Established new peer connection on Hub")
		case client := <-h.unsubscribe:
			// removes client connection status from pool
			if _, ok := h.connections[client]; ok {
				delete(h.connections, client)
				close(client.sink)
			}
			log.Println("[Socket] Disconnected peer connection on Hub")
		case message := <-h.broadcast:
			// broadcast message to all clients in poll
			log.Println("[Socket] Recieved message on Hub broadcast channel")
			for client := range h.connections {
				log.Println("[Socket] Broadcasting message on Hub")
				select {
				case client.sink <- message:
					log.Printf("[Socket] Message broadcasted to %v peer on Hub", client)
				default:
					close(client.sink)
					delete(h.connections, client)
				}
			}
		}
	}
}
