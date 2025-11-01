package hub

import (
	"chat-system/models"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	Username string
	Conn     *websocket.Conn
	Send     chan models.Message
	Hub      *Hub
	Room     *Room
}

func NewClient(conn *websocket.Conn, hub *Hub, room *Room, username string) *Client {
	return &Client{
		ID:       uuid.New().String(),
		Username: username,
		Conn:     conn,
		Send:     make(chan models.Message),
		Hub:      hub,
		Room:     room,
	}
}

func (c *Client) Read() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		msg := models.BuildMessage(models.ClientInfo{
			ID:       c.ID,
			Username: c.Username,
		}, c.Room.Name, models.CommandMessage, string(p))

		c.parseMessage(msg)
	}
}

func (c *Client) Write() {
	for msg := range c.Send {
		err := c.Conn.WriteJSON(msg)
		log.Printf("Message sent to client %s: %+v\n", c.ID, msg)

		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
	}
}

func (c *Client) parseMessage(messageInfo models.Message) {
	message := messageInfo.Text
	if strings.HasPrefix(message, "/") {
		log.Println("Parsing command:", message)

		parts := strings.SplitN(message, " ", 2)
		command := strings.ToLower(parts[0])

		switch command {
		case "/users":
			userList := c.Room.ListClients()
			if len(userList) == 0 {
				c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandListUsers, "No users in room")
				return
			}
			c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandListUsers, "Users in room: "+strings.Join(userList, ", "))
		case "/rooms":
			roomList := c.Hub.ListRooms()
			if len(roomList) == 0 {
				c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandListRooms, "No rooms available")
				return
			}
			c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandListRooms, "Rooms: "+strings.Join(roomList, ", "))
		case "/join":
			if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
				c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandMessage, "Usage: /join <room_name>")
				return
			}
			roomName := parts[1]
			c.Room.Unregister <- c

			newRoom := c.Hub.GetOrCreateRoom(roomName)
			c.Room = newRoom
			newRoom.Register <- c
		default:
			c.Send <- models.BuildMessage(models.ClientInfo{ID: c.ID, Username: c.Username}, c.Room.Name, models.CommandMessage, "Unknown command: "+command)
		}
	} else {
		c.Room.Broadcast <- messageInfo
		log.Printf("Message sent from client %s: %+v\n", c.ID, messageInfo)
	}
}
