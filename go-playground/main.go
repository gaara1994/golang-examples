package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	fmt.Println(validate)

}
