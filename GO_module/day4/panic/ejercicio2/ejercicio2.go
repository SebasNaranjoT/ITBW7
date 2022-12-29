package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ingresando al archivo de customers")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Finalizó la ejecución del programa")
	}()
	customers, err := os.ReadFile("customers.txt")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Se abrió con exito el archivo")
		fmt.Println(string(customers))
	}
}
