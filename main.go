package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`
}

func main() {

	var v = validator.New()

	user := &User{
		FirstName: "Badger",
		LastName:  "",
		Age:       99,
	}
	if err := v.Struct(user); err != nil {
		fmt.Println("エラーあり")
	}

}
