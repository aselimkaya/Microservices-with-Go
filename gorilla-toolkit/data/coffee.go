package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Coffee struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type CoffeeList []*Coffee

func (coffee *Coffee) ConvertFromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(coffee)
}

func (coffees *CoffeeList) ConvertToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(coffees)
}

func GetCoffeeList() CoffeeList {
	return coffeeList
}

func AddCoffee(coffee *Coffee) {
	coffee.ID = getNextID()
	coffeeList = append(coffeeList, coffee)
}

func UpdateCoffee(id int, coffee *Coffee) error {
	_, position, err := findByID(id)

	if err != nil {
		return err
	}

	coffee.ID = id

	coffeeList[position] = coffee

	return nil
}

func getNextID() int {
	return coffeeList[len(coffeeList)-1].ID + 1
}

var ErrorCoffeeNotFound = fmt.Errorf("Kahve bulunamadı")

func findByID(id int) (*Coffee, int, error) {
	for index, coffee := range coffeeList {
		if coffee.ID == id {
			return coffee, index, nil
		}
	}
	return nil, -1, ErrorCoffeeNotFound
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
