package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aselimkaya/microservices/gorilla-toolkit/handlers"
	"github.com/gorilla/mux"
)

func main() {
	coffeeLogger := log.New(os.Stdout, "coffee-shop-api", log.LstdFlags)
	coffeeHandler := handlers.NewCoffeeHandler(coffeeLogger)

	router := mux.NewRouter()

	//Subrouter kullanımı
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/coffees", coffeeHandler.GetAll) //Bu şekilde çağırabilmemiz için fonksiyon GetAll(responseWriter http.ResponseWriter, request *http.Request) şeklinde olmalı

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/coffee", coffeeHandler.AddCoffee)
	postRouter.Use(coffeeHandler.MiddlewareValidateCoffee)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/coffee/{id:[0-9]+}", coffeeHandler.UpdateCoffee)
	putRouter.Use(coffeeHandler.MiddlewareValidateCoffee)

	server := &http.Server{
		Addr:        ":9999",
		Handler:     router,
		IdleTimeout: 60 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			coffeeLogger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	coffeeLogger.Println("Sonlandırma isteği geldi.", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
