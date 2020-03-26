package data

import (
	"encoding/json"
	"io"
)

type Coffee struct {
	//Tag ile birlikte JSON gösterimindeki isimleri değişitiriyoruz
	//API çağrısında görmek istemediğimiz alanlar için `json:"-"`
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type CoffeeList []*Coffee

func (coffees *CoffeeList) ConvertToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(coffees)
}

//Getter fonksiyon
func GetCoffeeList() CoffeeList {
	return coffeeList
}

var coffeeList = CoffeeList{
	&Coffee{
		ID:    1,
		Name:  "Filtre Kahve",
		Price: 7.0,
	},
	&Coffee{
		ID:    2,
		Name:  "Espresso",
		Price: 5.5,
	},
	&Coffee{
		ID:    3,
		Name:  "Latte Macchiato",
		Price: 10.0,
	},
	&Coffee{
		ID:    4,
		Name:  "Espresso Macchiato",
		Price: 8.0,
	},
	&Coffee{
		ID:    5,
		Name:  "Latte",
		Price: 9.0,
	},
	&Coffee{
		ID:    6,
		Name:  "Flat White",
		Price: 11.5,
	},
}
