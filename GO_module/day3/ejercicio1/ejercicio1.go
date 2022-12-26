package main

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {

	for _, product := range Products {
		fmt.Printf("%+v\n", product)
	}
}

func (p Product) GetById(id int) (resultado Product) {
	for _, product := range Products {
		if product.Id == id {
			resultado = product
			return resultado
		} else {
			continue
		}
	}
	return
}

func main() {
	producto1 := Product{
		Id:          123,
		Name:        "Cebolla",
		Price:       2300.0,
		Description: "Cebolla cabezona",
		Category:    "Verduras",
	}

	producto2 := Product{
		Id:          124,
		Name:        "Tomate",
		Price:       1670.0,
		Description: "Tomate de aliño",
		Category:    "Verduras",
	}
	producto2.Save()
	producto1.Save()
	Product.GetAll(Product{})
	//producto1.GetAll()
	fmt.Printf("%+v\n", Product.GetById(Product{}, 123))
}
