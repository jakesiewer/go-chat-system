package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID   string     `json:"id"`
	From ClientInfo `json:"from"`
	Text string     `json:"text"`
	Room string     `json:"room"`
	Type Command    `json:"type"`
	Time time.Time  `json:"time"`
}

type ClientInfo struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type Command string

const (
	CommandMessage   Command = "message"
	CommandJoin      Command = "join"
	CommandLeave     Command = "leave"
	CommandListUsers Command = "list-users"
	CommandListRooms Command = "list-rooms"
)

func BuildMessage(from ClientInfo, room string, msgType Command, text string) Message {
	return Message{
		ID:   uuid.New().String(),
		From: from,
		Room: room,
		Type: msgType,
		Text: text,
		Time: time.Now(),
	}
}
