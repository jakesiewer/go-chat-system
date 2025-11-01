package hub

import (
	"chat-system/models"
	"log"

	"github.com/google/uuid"
)

type Room struct {
	ID         string
	Name       string
	Clients    map[*Client]bool
	Broadcast  chan models.Message
	Register   chan *Client
	Unregister chan *Client
}

func NewRoom(name string) *Room {
	log.Printf("Creating new room: %s", name)
	return &Room{
		ID:         uuid.New().String(),
		Name:       name,
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan models.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (room *Room) Run(h *Hub) {
	defer func() {
		log.Printf("Room %s is shutting down", room.Name)
		delete(h.Rooms, room.Name)
	}()

	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
		case client := <-room.Unregister:
			delete(room.Clients, client)
		case msg := <-room.Broadcast:
			for client := range room.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(room.Clients, client)
				}
			}
		}
	}
}

func (room *Room) ListClients() []string {
	clients := make([]string, 0, len(room.Clients))
	for client := range room.Clients {
		clients = append(clients, client.Username)
	}
	return clients
}
