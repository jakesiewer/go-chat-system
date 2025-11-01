package models

import "time"

type Message struct {
	ID   string     `json:"id,omitempty"`
	From ClientInfo `json:"from,omitempty"`
	Text string     `json:"text,omitempty"`
	Room string     `json:"room,omitempty"`
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
		From: from,
		Room: room,
		Type: msgType,
		Text: text,
		Time: time.Now(),
	}
}
