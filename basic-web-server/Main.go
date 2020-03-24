package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) { //özellikle handle etmediğim pathler için default handler görevi görüyor. Örn /xyz isteği bile yapılsa bu handler çalışıyor
		log.Println("Welcome!")

		d, err := ioutil.ReadAll(request.Body)

		if err != nil {
			http.Error(responseWriter, "Bir şeyler ters gitti!", http.StatusBadRequest)
			/* Alternatif olarak bu şekilde de yapılabilir
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte("Bir şeyler ters gitti!"))*/
			return //http.Error kodun çalışmasını sonlandırmıyor, o nedenle return kullanmam gerekiyor
		}

		fmt.Fprintf(responseWriter, "Request Body data is %s", d)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) { //bye path'ini handle ediyor
		log.Println("Bye!")
	})

	http.ListenAndServe(":9090", nil) //Bütün IP adreslerinden gelen bağlantılar dinleniyor
}
