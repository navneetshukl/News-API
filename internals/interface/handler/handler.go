
package handler

import (
	"encoding/json"
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
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.connections = append(ws.connections, conn)
}

func (ws *WebSocket) Remove(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for i, c := range ws.connections {
		if c == conn {
			// Remove the connection by slicing
			ws.connections = append(ws.connections[:i], ws.connections[i+1:]...)
			break
		}
	}
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

	go ws.BroadCast(conn)

	// Handle incoming messages from the client
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

	// Remove the connection when done
	ws.Remove(conn)
}

func (ws *WebSocket) BroadCast(conn *websocket.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Send a message to the client at regular intervals
	for {
		select {
		case <-ticker.C:
			n1, _ := ws.NewsUseCase.GetFirstNews()
			n2, _ := ws.NewsUseCase.GetSecondNews()
			n3, _ := ws.NewsUseCase.GetThirdNews()
			allArticles := []news.Article{}

			allArticles = append(allArticles, n1.Articles...)
			allArticles = append(allArticles, n2.Articles...)
			allArticles = append(allArticles, n3.Articles...)

			allNews := news.AllNews{}
			allNews.Articles = allArticles
			ws.mu.Lock()
			allNews.Connected_Users = len(ws.connections)
			ws.mu.Unlock()

			msg, _ := json.Marshal(allNews)
			err := conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return // Stop if there is an error
			}
		}
	}
}
