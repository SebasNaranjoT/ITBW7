package main

import (
	"fmt"
)

func main(){
	Letras("Mazaguaro")
}

func Letras(palabra string) {
	var cantLetras = len(palabra)
	fmt.Println("La cantidad de letras es: ", cantLetras)
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}
