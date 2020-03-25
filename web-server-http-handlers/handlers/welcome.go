package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Welcome struct {
	welcomeLogger *log.Logger
}

func NewWelcomeHandler(welcomeLogger *log.Logger) *Welcome { //Dependency injection
	return &Welcome{welcomeLogger}
}

func (welcome *Welcome) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	welcome.welcomeLogger.Println("Hoş geldiniz!")

	d, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(responseWriter, "Bir şeyler ters gitti!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(responseWriter, "Hoş geldin %s\n", d)
}
