package main

import "fmt"

func main() {
	fmt.Println(promedio(4.8, 2.5, 2))
}

func promedio(notes ...float64) (promedio float64) {
	var resultado float64
	for _, note := range notes {
		resultado += note
	}
	promedio = resultado / float64(len(notes))
	return promedio

}
