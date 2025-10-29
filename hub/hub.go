package hub

import (
	"chat-system/models"
	"log"
)

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan models.Message
}

func NewHub() *Hub {
	log.Println("Creating new Hub")
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan models.Message),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			hub.registerClient(client)
		case client := <-hub.Unregister:
			hub.unregisterClient(client)
		case message := <-hub.Broadcast:
			hub.broadcastMessage(message)
		}
	}
}

func (hub *Hub) registerClient(client *Client) {
	client.Room.Register <- client
}

func (hub *Hub) unregisterClient(client *Client) {
	client.Room.Unregister <- client
}

func (hub *Hub) broadcastMessage(message models.Message) {
	for _, room := range hub.Rooms {
		for client := range room.Clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(room.Clients, client)
			}
		}
	}
}

func (hub *Hub) ListRooms() []string {
	var roomList []string
	for roomName := range hub.Rooms {
		roomList = append(roomList, roomName)
	}
	return roomList
}

func (hub *Hub) GetOrCreateRoom(roomName string) *Room {
	room := hub.Rooms[roomName]
	if room == nil {
		room = NewRoom(roomName)
		hub.Rooms[roomName] = room
		go room.Run(hub)
	}
	return room
}
