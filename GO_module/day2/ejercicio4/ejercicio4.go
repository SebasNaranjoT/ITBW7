package main

import (
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func minFunc(valores ...float64) (minValue float64) {
	minValue = valores[0]

	for _, valor := range valores {
		if valor < minValue {
			minValue = valor
		}
	}
	return minValue
}

func maxFunc(valores ...float64) (maxValue float64) {
	maxValue = valores[0]

	for _, valor := range valores {
		if valor > maxValue {
			maxValue = valor
		}
	}
	return maxValue
}

func avgFunc(notes ...float64) (average float64) {
	var resultado float64
	for _, note := range notes {
		resultado += note
	}
	average = resultado / float64(len(notes))
	return average

}

func wrongOperation(notes ...float64) (resultado float64) {
	resultado = 0.0
	fmt.Println("Ingrese una operación válida")
	return
}

func operationFunc(operation string) func(valores ...float64) float64 {

	switch operation {

	case minimum:
		return minFunc
	case maximum:
		return maxFunc
	case average:
		return avgFunc
	}
	return wrongOperation
}

func main() {

	resultado := operationFunc("averae")
	fmt.Println(resultado(3.0, 3.0, 3.0))
}
