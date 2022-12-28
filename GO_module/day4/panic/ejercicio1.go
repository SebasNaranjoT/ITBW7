package main

import (
	"fmt"
	"os"
)

func leerArchivo() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err1 := os.Open("customers.txt")

	if err1 != nil {
		panic(err1)
	}
}

func main() {

	fmt.Println("Inicia ejecución")

	leerArchivo()

	fmt.Println("Fin de la ejecución")
}
