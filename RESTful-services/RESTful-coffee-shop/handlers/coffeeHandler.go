package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/aselimkaya/microservices/RESTful-services/RESTful-coffee-shop/data"
)

type coffeeList struct {
	coffeeLogger *log.Logger
}

func NewCoffeeHandler(l *log.Logger) *coffeeList {
	return &coffeeList{l}
}

func (coffee *coffeeList) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		coffee.getAll(responseWriter)
		return
	} else if request.Method == http.MethodPost {
		coffee.addCoffee(responseWriter, request)
		return
	} else if request.Method == http.MethodPut {
		//URL'den ID'yi elde etmemiz gerekecek. Hangi objeyi güncelleyeceğimizi bu şekilde öğreneceğiz.
		matchGroup := regexp.MustCompile(`/([0-9]+)`).FindAllStringSubmatch(request.URL.Path, -1)

		if len(matchGroup) != 1 && len(matchGroup[0]) != 2 {
			coffee.coffeeLogger.Println("Geçersiz URI")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := matchGroup[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			coffee.coffeeLogger.Println("ID dönüşümü yapılamadı")
			http.Error(responseWriter, "Invalid URI", http.StatusBadRequest)
			return
		}

		coffee.updateCoffee(id, responseWriter, request)

		return
	}

	//else durumu
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (coffee *coffeeList) getAll(responseWriter http.ResponseWriter) {
	coffee.coffeeLogger.Println("HTTP GET isteği")

	listOfCoffees := data.GetCoffeeList()

	err := listOfCoffees.ConvertToJSON(responseWriter)

	if err != nil {
		http.Error(responseWriter, "Veriler çağrılırken bir hata oluştu!", http.StatusInternalServerError)
	}
}

func (coffee *coffeeList) addCoffee(responseWriter http.ResponseWriter, request *http.Request) {
	coffee.coffeeLogger.Println("HTTP POST isteği")

	newCoffee := &data.Coffee{}

	err := newCoffee.ConvertFromJSON(request.Body)

	if err != nil {
		http.Error(responseWriter, "Veri işlenirken bir hata oluştu!", http.StatusBadRequest)
	}

	data.AddCoffee(newCoffee)
}

func (coffee *coffeeList) updateCoffee(id int, responseWriter http.ResponseWriter, request *http.Request) {
	coffee.coffeeLogger.Println("HTTP PUT isteği")

	newCoffee := &data.Coffee{}

	err := newCoffee.ConvertFromJSON(request.Body)

	if err != nil {
		http.Error(responseWriter, "Veri işlenirken bir hata oluştu!", http.StatusBadRequest)
	}

	err = data.UpdateCoffee(id, newCoffee)

	if err == data.ErrorCoffeeNotFound {
		http.Error(responseWriter, "Kahve bulunamadı!", http.StatusNotFound)
	}

	if err != nil {
		http.Error(responseWriter, "Kahve bulunamadı!", http.StatusInternalServerError)
	}
}
