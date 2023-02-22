package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la funcion: ", cadena)

	modificarValor(&cadena) // Referencia
	//modificarValor(cadena) Copia

	fmt.Println("Despues de la cadena: ", cadena)
}
