package websocket

import (
	"encoding/json"

	"github.com/go-pg/pg/v9"
	"github.com/justcompile/midgard/common"
	"github.com/justcompile/midgard/common/events"
	log "github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	eventListener *pg.Listener

	done chan bool
}

func (h *Hub) Run() {
	internalEvents := h.eventListener.Channel()

	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.onRegister(client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case notification := <-internalEvents:
			log.Debug("Received internal event: ", notification.Payload)
			go func() { h.broadcast <- []byte(notification.Payload) }()
		case <-h.done:
			h.eventListener.Close()

			for client := range h.clients {
				close(client.send)
				delete(h.clients, client)
			}

			return
		}
	}
}

// Shutdown informs the bug to close any open client connections
func (h *Hub) Shutdown() {
	h.done <- true
}

func (h *Hub) onRegister(client *Client) {
	workers := common.WorkerRegistry.GetConnectedWorkers()

	out, _ := json.Marshal(workers)

	client.send <- out
}

func NewHub() *Hub {
	return &Hub{
		broadcast:     make(chan []byte),
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		clients:       make(map[*Client]bool),
		eventListener: events.Subscribe("worker_connected", "worker_disconnected"),
	}
}
