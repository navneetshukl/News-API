package routes

import (
	"net/http"
	"news-api/internals/interface/handler"

	"github.com/gorilla/mux"
)

func Routes(WsHandler *handler.WebSocket) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ws", WsHandler.HandleConnection)

	return r
}
