// package handler

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"news-api/internals/core/news"
// 	"sync"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// type WebSocket struct {
// 	connections []*websocket.Conn
// 	mu          sync.Mutex
// 	NewsUseCase news.NewsUseCase
// }

// func NewWebsocket(news news.NewsUseCase) *WebSocket {
// 	return &WebSocket{
// 		NewsUseCase: news,
// 	}
// }

// func (ws *WebSocket) Add(conn *websocket.Conn) {
// 	ws.mu.Lock()
// 	defer ws.mu.Unlock()
// 	ws.connections = append(ws.connections, conn)
// }

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true // Allow connections from any origin
// 	},
// }

// func (ws *WebSocket) HandleConnection(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade the HTTP connection to a WebSocket connection
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error during connection upgrade:", err)
// 		return
// 	}
// 	ws.Add(conn)
// 	defer conn.Close()
// 	fmt.Println("Client connected")

// 	log.Println("Total number of connections is", len(ws.connections))

// 	for {
// 		go ws.handleMessages()
// 		time.Sleep(1*time.Minute)

// 	}
// }

// func (ws *WebSocket) handleMessages() {
// 	for {

// 		n1, _ := ws.NewsUseCase.GetFirstNews()
// 		bytes, err := json.Marshal(n1)
// 		if err != nil {
// 			log.Println("error in converting to byte ", err)
// 			continue
// 		}

// 		msg := bytes
// 		ws.broadcast(1, msg)
// 		time.Sleep(1 * time.Minute) // Add sleep to avoid flooding
// 	}
// }

// func (ws *WebSocket) broadcast(messageType int, msg []byte) {
// 	ws.mu.Lock()
// 	defer ws.mu.Unlock()

// 	for _, conn := range ws.connections {
// 		// Write to each connection one at a time
// 		err := conn.WriteMessage(messageType, msg)
// 		if err != nil {
// 			log.Println("Error writing message:", err)
// 			return
// 		}
// 	}
// }

package handler

import (
	"encoding/json"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	ws.Add(conn)
	defer func() {
		log.Println("Client disconnected")
		ws.Remove(conn) // Ensure to remove the connection on disconnect
		conn.Close()
	}()
	fmt.Println("Client connected")

	log.Println("Total number of connections is", len(ws.connections))

	// Start handling messages for this connection
	go ws.handleMessages(conn)

	// Keep reading messages from this connection
	
}

func (ws *WebSocket) Remove(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for i, c := range ws.connections {
		if c == conn {
			ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
			break
		}
	}
}

func (ws *WebSocket) handleMessages(conn *websocket.Conn) {
	for {
		n1, err := ws.NewsUseCase.GetFirstNews()
		if err != nil {
			log.Println("Error getting news:", err)
			continue
		}

		bytes, err := json.Marshal(n1)
		if err != nil {
			log.Println("Error converting to byte:", err)
			continue
		}

		msg := bytes
		ws.broadcast(websocket.TextMessage, msg)

		time.Sleep(1 * time.Minute) // Add sleep to avoid flooding
	}
}

func (ws *WebSocket) broadcast(messageType int, msg []byte) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	for i := 0; i < len(ws.connections); i++ {
		conn := ws.connections[i]

		// Write to each connection one at a time
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println("Error writing message:", err)
			conn.Close() // Close the connection if there's an error
			// Remove the closed connection
			ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
			i-- // Decrement i to check the new connection at this index
		}
	}
}
