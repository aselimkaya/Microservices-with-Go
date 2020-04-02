package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

type User struct {
	SKU       string `json:"sku" validate:"required,sku"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
	Email     string `validate:"required,email"`
}

func (user *User) Validate() error {
	validate := validator.New()

	//Kendi validation fonksiyonlarımızı da yazabiliyoruz
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(user)
}

func validateSKU(fl validator.FieldLevel) bool {
	//SKU formatı abc-defg-hijk şeklinde
	rgx := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := rgx.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func NewUser() *User {
	return &User{}
}
