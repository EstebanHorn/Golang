package main

import "fmt"

func saludo(nombre string) {
	fmt.Println("Hola", nombre)
}

func suma(a, b int) int {
	return a + b
}

func main() {
	var nombre string
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	saludo(nombre)

	var a, b int
	fmt.Println("Ingrese un numero: ")
	fmt.Scanln(&a)
	fmt.Println("Ingrese otro numero: ")
	fmt.Scanln(&b)
	res := suma(a, b)
	fmt.Println("la suma es", res)
}
