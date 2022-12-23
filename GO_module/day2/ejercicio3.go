package main

import "fmt"

func calculoHoras(minutos int) (horas float64) {
	horas = float64(minutos) / 60.0
	return horas
}

func calcularSalario(categoria string, minutos int) (salario float64) {
	switch categoria {
	case "A":
		salario = (3000 * calculoHoras(minutos)) * (1 + 0.5)
		return salario
	case "B":
		salario = (1500 * calculoHoras(minutos)) * (1 + 0.2)
		return salario
	case "C":
		salario = (1000 * calculoHoras(minutos))
		return salario
	}
	return
}

func main() {
	var (
		categoria string
		minutos   int
	)

	categoria = "A"
	minutos = 48
	fmt.Println("Su salario es: ", calcularSalario(categoria, minutos))
}
