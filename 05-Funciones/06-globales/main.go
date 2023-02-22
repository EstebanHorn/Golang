package main

import "fmt"

var mensaje string

func funcion1() {
	mensaje = "mensaje funcion 1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "mensaje funcion 2"
	fmt.Println(mensaje)

}

func main() {

	mensaje = "mensaje principal"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

	fmt.Println(mensaje)

}
