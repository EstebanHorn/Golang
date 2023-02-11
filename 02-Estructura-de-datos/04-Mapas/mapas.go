package main

import (
	"fmt"
)

func main() {
	dias := make(map[int]string)

	fmt.Println(dias)

	//agregar
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"
	fmt.Println(dias)

	dias[1] = "Monday"
	fmt.Println(dias)

	delete(dias, 1)
	fmt.Println(dias)

	fmt.Println(dias, len(dias))

	//nuevo mapa
	estudaiantes := make(map[string][]int)

	estudaiantes["Esteban"] = []int{1, 2, 3}
	estudaiantes["Samuel"] = []int{10, 9, 8}

	fmt.Println(estudaiantes)
	fmt.Println(estudaiantes["Esteban"])
	fmt.Println(estudaiantes["Esteban"][1])

}
