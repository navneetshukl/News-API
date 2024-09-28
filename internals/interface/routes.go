package routes

import (
	"net/http"
	"news-api/internals/interface/handler"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes(WsHandler *handler.WebSocket) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ws", WsHandler.HandleConnection)

	// Create a CORS handler
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Update this to specific origins in production
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	return corsHandler.Handler(r) // No need for type assertion
}
