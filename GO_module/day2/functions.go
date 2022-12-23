package main

import "fmt"

const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
)

func main() {
	fmt.Println(operacionAritmetica(6, 2, Suma))
	fmt.Println(operacionAritmetica(6, 2, Resta))
	fmt.Println(operacionAritmetica(6, 2, Divis))
	fmt.Println(operacionAritmetica(6, 2, Multip))
}

func operacionAritmetica(a, b float64, operacion string) float64 {
	switch operacion {
	case Suma:
		return a + b
	case Resta:
		return a - b
	case Divis:
		return a / b
	case Multip:
		return a * b
	}
	return 0
}
