package main

import (
	"fmt"
)

func main() {
	fmt.Println("ingrese su edad")
	var age int
	fmt.Scanln(&age)

	fmt.Println("Está trabajando?")
	var empleo bool
	fmt.Scanln(&empleo)

	fmt.Println("Cual es su salario?")
	var salario int
	fmt.Scanln(&salario)

	fmt.Println("cuantos años lleva trabajando?")
	var años int
	fmt.Scanln(&años)

	mensaje(age, empleo, salario, años)
}

func mensaje(edad int, empleo bool, sueldo int, años int) {
	if (edad > 22 && empleo) && años > 1 {
		if sueldo > 100000 {
			fmt.Println("Tu prestamo ha sido aprobado sin intereses")
		} else {
			fmt.Println("Debes pagar un interes de 18%")
		}
	} else {
		fmt.Println("Tu prestamo no ha sido aprobado")
	}
}
