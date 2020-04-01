package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aselimkaya/microservices/gorilla-toolkit/websocket/handler"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "websocket", log.LstdFlags)
	h := handler.NewWebsocketHandler(l)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", h.WebsocketEndpoint)

	log.Fatal(http.ListenAndServe(":8888", router))
}
