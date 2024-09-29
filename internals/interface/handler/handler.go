package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news-api/internals/core/news"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin (not recommended for production)
	},
}

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
	ws.connections = append(ws.connections, conn)
}

func (ws *WebSocket) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()
	ws.Add(conn)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	//	cnt := 0

	// Send a message to the client at regular intervals
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			str := fmt.Sprintf("%s,%d", "Hello from the server!", cnt)
	// 			message := map[string]string{"message": str}
	// 			msg, _ := json.Marshal(message)
	// 			err := conn.WriteMessage(websocket.TextMessage, msg)
	// 			if err != nil {
	// 				return // Stop if there is an error
	// 			}
	// 			cnt++
	// 		}
	// 	}
	// }()

	go ws.BroadCast(conn)

	//Handle incoming messages from the client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Echo the received message back to the client
		err = conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

func (ws *WebSocket) BroadCast(conn *websocket.Conn) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	cnt := 0

	// Send a message to the client at regular intervals
	for {
		select {
		case <-ticker.C:
			str := fmt.Sprintf("%s,%d", "Hello from the server!", cnt)
			message := map[string]string{"message": str}
			msg, _ := json.Marshal(message)
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return // Stop if there is an error
			}
			cnt++
		}
	}

}
