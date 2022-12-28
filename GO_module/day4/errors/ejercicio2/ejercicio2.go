package main

import (
	"errors"
	"fmt"
)

type customError struct {
	msg string
}

func (c customError) Error() string {
	return "Error: El salario es menor a 10000"
}

func NewCustom(message string) error {
	return &customError{msg: message}
}

var otherCustom = NewCustom("Otro mensaje de error custom")

func main() {

	var salary int = 1000
	err := validacionSalario(salary)
	if err != nil {
		if errors.Is(err, otherCustom) {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("Debe pagar Impuesto")
	}
}

func validacionSalario(salary int) error {
	if salary <= 10000 {
		return otherCustom
	}

	return nil
}
