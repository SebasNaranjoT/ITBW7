package main

import "fmt"

type customError struct {
	//
}

func (e *customError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	var salary int = 160000

	if salary < 150000 {
		e := customError{}
		fmt.Println(e.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}

}


Jean Pooll Marambio Cayo11:40
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

var (
	ErrOther = NewCustom("other error from custom")
)

func main() {
	var salary int = 1000
	err := validarSalario(salary)
	if err != nil {
		if errors.Is(err, ErrOther) {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Debe pagar Impuesto")

}
	}
}

func validarSalario(salary int) error {
	if salary < 10000 {
		return ErrOther
	}

	return nil
}