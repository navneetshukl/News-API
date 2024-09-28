package routes

import (
	"news-api/internals/interface/handler"

	"github.com/gorilla/mux"
)

func Routes(WsHandler *handler.WebSocket) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/ws", WsHandler.HandleConnection)

	return r

}
