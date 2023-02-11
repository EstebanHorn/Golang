package main

import "fmt"

func main() {
	hola := "hola"
	mundo := "mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Esteban"
	edad := 19

	fmt.Printf("Hola, %s tu edad es %d \n", nombre, edad)

	fmt.Printf("Hola, %v tu edad es %v \n", nombre, edad)

	msj := fmt.Sprintf("Hola, %v tu edad es %v", nombre, edad)

	fmt.Println(msj)

	fmt.Printf("nombre: %T \n", nombre)

	fmt.Print("Ingerese su nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("otro nombre ", nombre)
}
