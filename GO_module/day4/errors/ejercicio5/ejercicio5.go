package main

import (
	"errors"
	"fmt"
)

func calcularSalario(hours int, costo float64) (err error, salary float64) {
	if hours < 80 {
		return errors.New("Error: el trabajador no puede haber trabajado menos de 80 Hrs"), salary
	} else {
		salary = costo * float64(hours)
		if salary >= 150000.0 {
			return err, salary * (1 - 0.1)
		}
	}
	return
}

func main() {
	var (
		costPerHour float64
		hours       int
	)

	costPerHour = 2000
	hours = 100

	fmt.Println(calcularSalario(hours, costPerHour))

}
