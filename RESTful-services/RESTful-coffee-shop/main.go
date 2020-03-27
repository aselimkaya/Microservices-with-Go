package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aselimkaya/microservices/RESTful-services/RESTful-coffee-shop/handlers"
)

func main() {
	coffeeLogger := log.New(os.Stdout, "coffee-shop-api", log.LstdFlags)
	coffeeHandler := handlers.NewCoffeeHandler(coffeeLogger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", coffeeHandler)

	server := &http.Server{
		Addr:        ":9999",
		Handler:     serveMux,
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
