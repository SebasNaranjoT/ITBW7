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
