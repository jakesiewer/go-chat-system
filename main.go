package main

import (
	"chat-system/hub"
	"chat-system/server"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Distributed Chat App v0.01")

	h := hub.NewHub()
	go h.Run()

	h.GetOrCreateRoom("general")

	setupRoutes(h)
	http.ListenAndServe(":8080", nil)
}

func setupRoutes(h *hub.Hub) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(h, w, r)
	})
}

func serveWs(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := server.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		return
	}

	username, room := parseRequest(r, h)
	client := hub.NewClient(conn, h, room, username)

	h.Register <- client

	go client.Read()
	go client.Write()
}

func parseRequest(r *http.Request, h *hub.Hub) (string, *hub.Room) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "anonymous"
	}

	roomName := r.URL.Query().Get("room")
	var room *hub.Room
	if roomName == "" {
		room = h.Rooms["general"]
		return username, room
	}

	room = h.GetOrCreateRoom(roomName)
	return username, room
}
