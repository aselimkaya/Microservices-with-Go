package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	goodbyeLogger *log.Logger
}

func NewGoodbyeHandler(goodbyeLogger *log.Logger) *Goodbye {
	return &Goodbye{goodbyeLogger}
}

func (goodbye *Goodbye) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	goodbye.goodbyeLogger.Println("Güle güle!")

	d, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(responseWriter, "Bir şeyler ters gitti!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(responseWriter, "Güle güle %s\n", d)
}
