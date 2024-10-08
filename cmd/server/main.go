package main

import (
	"log"
	"net/http"
	routes "news-api/internals/interface"
	"news-api/internals/interface/handler"
	"news-api/internals/usecase/news"

	"github.com/joho/godotenv"
)

func main() {

	err:=godotenv.Load(".env")
	if err!=nil{
		log.Println("error in loading .env ",err)
		return
	}

	newsUsecase := news.NewNewsUsecase()

	wsHandler := handler.NewWebsocket(newsUsecase)
	router := routes.Routes(wsHandler)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("error in connecting to websocket ", err)
		return
	}

}
