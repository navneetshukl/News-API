package handler

import (
	"fmt"
	"log"
	"net/http"
	"news-api/internals/core/news"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	connections []*websocket.Conn
	mu          sync.Mutex
	NewsUseCase news.NewsUseCase
}

func NewWebsocket(news news.NewsUseCase) *WebSocket {
	return &WebSocket{
		NewsUseCase: news,
	}
}

func (ws *WebSocket) Add(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.connections = append(ws.connections, conn)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func (ws *WebSocket) HandleConnection(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	ws.Add(conn)
	defer conn.Close()
	fmt.Println("Client connected")

	log.Println("Total number of connections is", len(ws.connections))

	for {
		go ws.handleMessages()

	}
}

func (ws *WebSocket) handleMessages() {
	for {

		msg := []byte("Hi")
		ws.broadcast(1, msg)
		time.Sleep(1 * time.Minute) // Add sleep to avoid flooding
	}
}

func (ws *WebSocket) broadcast(messageType int, msg []byte) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for _, conn := range ws.connections {
		// Write to each connection one at a time
		err := conn.WriteMessage(messageType, msg)
		if err != nil {
			log.Println("Error writing message:", err)
		}
	}
}
