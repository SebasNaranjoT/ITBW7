package main

import "testing"

func TestCalculoImpuesto(t *testing.T) {
	//Arrange
	salario1 := 10000.0
	impuestoEsperado1 := 0.0
	salarioEsperado1 := 10000.0

	salario2 := 60000.0
	impuestoEsperado2 := 10200.0
	salarioEsperado2 := 49800.0

	salario3 := 160000.0
	impuestoEsperado3 := 4200.0
	salarioEsperado3 := 116800.0

	resultado1, impuesto1 := calculoImpuesto(salario1)
	resultado2, impuesto2 := calculoImpuesto(salario2)
	resultado3, impuesto3 := calculoImpuesto(salario3)

	if resultado1 != salarioEsperado1 || impuesto1 != impuestoEsperado1 {
		t.Fatalf("Resultado inesperado para el salario 1")
	}

	if resultado2 != salarioEsperado2 || impuesto2 != impuestoEsperado2 {
		t.Fatalf("Resultado inesperado para el salario 2")
	}

	if resultado3 != salarioEsperado3 || impuesto3 != impuestoEsperado3 {
		t.Fatalf("Resultado inesperado para el salario 3")
	}

}
