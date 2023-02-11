package main

import (
	"fmt"
)

func main() {

	//arrays
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 20
	numeros[1] = 30
	numeros[4] = 10
	fmt.Println(numeros)

	//arrays con valores
	nombres := [3]string{"Esteban", "Samuel", "Matias"}
	fmt.Println(nombres)

	colores := [...]string{
		"Rojo",
		"Azul",
		"Verde",
		"Rosa",
	}

	fmt.Println(colores, len(colores))

	//indices indefinidos
	monedas := [...]string{
		0: "dolar",
		1: "peso",
		3: "euro",
	}

	fmt.Println(monedas, len(monedas))

	//Subarray
	subMoneda := monedas[0:2]

	fmt.Println(subMoneda)
}
