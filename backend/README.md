# Go Chat System

This is a real-time chat application built with Go and WebSockets. It allows multiple clients to connect to chat rooms and exchange messages in real-time.

## Features

*   **Real-time messaging:** Send and receive messages instantly.
*   **Multiple chat rooms:** Create and join different chat rooms.
*   **List users:** See a list of users in the current room.
*   **List rooms:** See a list of all available rooms.
*   **Switch rooms:** Easily switch between chat rooms.

## Getting Started

### Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.15+ recommended)

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/jakesiewer/go-chat-system.git
    cd go-chat-system
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Run the server:**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

## Usage

To connect to the chat server, you can use a WebSocket client like [Simple WebSocket Client](https://chrome.google.com/webstore/detail/simple-websocket-client/pfdhoblngboilpfeibdedpjgfnlcodoo?hl=en) for Chrome.

1.  **Connection URL:**
    ```
    ws://localhost:8080/ws?username=your-username&room=your-room
    ```
    *   Replace `your-username` with your desired username. If not provided, it defaults to "anonymous".
    *   Replace `your-room` with the name of the room you want to join. If not provided, you will join the "general" room.

2.  **Chat Commands:**
    *   `/users`: Lists all users in the current room.
    *   `/rooms`: Lists all available chat rooms.
    *   `/join <room_name>`: Switches to the specified chat room.

## Project Structure

```
.
├── hub/
│   ├── client.go       # WebSocket client logic
│   ├── hub.go          # Manages rooms and clients
│   └── room.go         # Chat room logic
├── main.go             # Application entry point
├── models/
│   └── message.go      # Defines the chat message structure
└── server/
    └── websocket_handler.go # WebSocket connection handler
```
