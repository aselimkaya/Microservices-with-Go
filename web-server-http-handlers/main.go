package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aselimkaya/microservices/web-server-http-handlers/handlers"
)

func main() {
	logger := log.New(os.Stdout, "test-api", log.LstdFlags)
	welcomeHandler := handlers.NewWelcomeHandler(logger)
	goodbyeHandler := handlers.NewGoodbyeHandler(logger)

	serveMux := http.NewServeMux() //Default Serve Mux'u kullanmıyoruz.
	serveMux.Handle("/", welcomeHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	server := &http.Server{ //http.ListenAndServe yerine özelleştirebildiğimiz HTTP Sunucusu
		Addr:        ":9090",
		Handler:     serveMux, //nil yazdığımızda Default Serve Mux kullanılıyordu. Artık kendi tanımladığımız kullanılacak.
		IdleTimeout: 60 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	/*
		İşletim sistemi seviyesinde sinyalleri dinleyebildiğimiz bir kanal oluşturuyoruz.
		Interrupt veya Kill sinyali gelirse sunucu bağlantısını "Graceful" olarak sonlanırıyoruz.
		Graceful Shutdown: Yeni bağlantılara izin verilmiyor, mevcut bağlantıların işini bitirmesi için bekleniyor.
	*/
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	sig := <-signalChannel
	logger.Println("Sonlandırma isteği geldi.", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
