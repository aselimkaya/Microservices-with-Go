package test

import (
	"testing"

	"github.com/aselimkaya/microservices/validator/data"
)

func TestUserValidation(test *testing.T) {

	u := data.NewUser()
	u.FirstName = "Selim"
	u.LastName = "Kaya"
	u.Age = 99
	u.Email = "aselimkaya35@gmail.com"
	u.SKU = "abc-defg-hijk"

	err := u.Validate()

	if err != nil {
		test.Fatal(err)
	}
}
