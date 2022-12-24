package main

import "fmt"

func calculoImpuesto(salario float64) (salarioNeto, impuesto float64) {
	if salario > 50000 && salario <= 150000 {
		impuesto = salario * 0.17
		salarioNeto = salario - impuesto
		return salarioNeto, impuesto
	} else if salario > 150000 {
		impuesto = salario * 0.27
		salarioNeto = salario - impuesto
		return salarioNeto, impuesto
	} else {
		impuesto = 0
		salarioNeto = salario
		return salarioNeto, impuesto
	}
}

func main() {
	fmt.Println("Ingrese su salario")
	var salario float64
	fmt.Scanln(&salario)

	neto, impuesto := calculoImpuesto(salario)
	fmt.Printf("Su salario neto es %v y pagó %v de impuestos", neto, impuesto)
}
