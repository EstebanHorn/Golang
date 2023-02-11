package main

import "fmt"

func main() {

	//Slicen
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	//Agrgar datos

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//Subslicen
	subSlicen := numeros[:2]

	numeros[0] = 7

	fmt.Println(numeros, len(numeros))
	fmt.Println(subSlicen, len(subSlicen))

	//capacidad

	meses := []string{
		"Enero",
		"Feberero",
		"Marzo",
	}

	fmt.Printf("len : %v, cap : %v, mem: %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	fmt.Printf("len : %v, cap : %v, mem: %p \n", len(meses), cap(meses), meses)

}
