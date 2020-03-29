package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/aselimkaya/microservices/gorilla-toolkit/data"
	"github.com/gorilla/mux"
)

type coffeeList struct {
	coffeeLogger *log.Logger
}

func NewCoffeeHandler(l *log.Logger) *coffeeList {
	return &coffeeList{l}
}

func (coffee *coffeeList) GetAll(responseWriter http.ResponseWriter, request *http.Request) {
	coffee.coffeeLogger.Println("HTTP GET isteği")

	listOfCoffees := data.GetCoffeeList()

	err := listOfCoffees.ConvertToJSON(responseWriter)

	if err != nil {
		http.Error(responseWriter, "Veriler çağrılırken bir hata oluştu!", http.StatusInternalServerError)
	}
}

func (coffee *coffeeList) AddCoffee(responseWriter http.ResponseWriter, request *http.Request) {
	coffee.coffeeLogger.Println("HTTP POST isteği")

	newCoffee := request.Context().Value(Key{}).(data.Coffee)

	data.AddCoffee(&newCoffee)
}

func (coffee *coffeeList) UpdateCoffee(responseWriter http.ResponseWriter, request *http.Request) {
	//URL parametrelerini mux.Vars() fonksiyonu ile kolaylıkla elde ediyoruz
	vars := mux.Vars(request)
	idString := vars["id"]
	id, convErr := strconv.Atoi(idString)

	if convErr != nil {
		http.Error(responseWriter, "Nümerik dönüşüm sırasında hata ile karşılaşıldı", http.StatusBadRequest)
	}

	coffee.coffeeLogger.Println("HTTP PUT isteği")

	newCoffee := request.Context().Value(Key{}).(data.Coffee)

	err := data.UpdateCoffee(id, &newCoffee)

	if err == data.ErrorCoffeeNotFound {
		http.Error(responseWriter, "Kahve bulunamadı!", http.StatusNotFound)
	}

	if err != nil {
		http.Error(responseWriter, "Kahve bulunamadı!", http.StatusInternalServerError)
	}
}

type Key struct{}

//JSON Serialize işleminde kod tekrarı yapıyorduk. Bunu önlemek için Middleware kullanıyoruz
func (coffee *coffeeList) MiddlewareValidateCoffee(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		newCoffee := &data.Coffee{}

		err := newCoffee.ConvertFromJSON(request.Body)
		if err != nil {
			http.Error(responseWriter, "Veri işlenirken bir hata oluştu!", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), Key{}, newCoffee)
		request = request.WithContext(ctx)

		// Middleware'da başka bir chain olabilir, o çağrılıyor
		next.ServeHTTP(responseWriter, request)
	})
}
