package handlers

import (
	"log"
	"net/http"

	"github.com/aselimkaya/microservices/RESTful-services/simple-coffee-shop/data"
)

type CoffeeList struct {
	coffeeLogger *log.Logger
}

//Dependency injection
func NewCoffeeHandler(l *log.Logger) *CoffeeList {
	return &CoffeeList{l}
}

func (coffee *CoffeeList) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		coffee.getList(responseWriter)
		return
	}

	//else durumu
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

//Encoder kullanarak JSON'a çevirme
func (coffee *CoffeeList) getList(responseWriter http.ResponseWriter) {
	listOfCoffees := data.GetCoffeeList()

	err := listOfCoffees.ConvertToJSON(responseWriter)

	if err != nil {
		http.Error(responseWriter, "Veri işlenirken bir hata oluştu!", http.StatusInternalServerError)
	}
}

/*
json.Marshal() fonksionu ile veriyi basit bir şekilde JSON formatına çeviriyoruz.
Fakat veriyi tutmak için yeterli buffer alanı olmayabilir, JSON dosyası çok büyük boyutlu olabilir.
Bu nedenle Encoder kullanmak daha mantıklı. Encoder io.Writer türünde bir parametre alıp dönüştürdüğü veriyi doğrudan bu Writer'a gönderip yazılmasını sağlıyor.
Üstelik daha hızlı.

func (coffee *CoffeeList) getCoffeeList(responseWriter http.ResponseWriter, request *http.Request) {
	listOfCoffees := data.GetCoffeeList()

	//Veriyi JSON'a dönüştürüyoruz
	//Marshal fonksiyonu dolaşıp bize byte dizisi olarak dönüyor
	data, err := json.Marshal(listOfCoffees)

	if err != nil {
		http.Error(responseWriter, "Veri işlenirken bir hata oluştu!", http.StatusInternalServerError)
	}

	responseWriter.Write(data)
}
*/
