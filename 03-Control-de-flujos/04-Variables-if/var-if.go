package main

import (
	"fmt"
)

func main() {

	if nombre, edad := "Esteban", 19; nombre == "Esteban" {
		fmt.Println("Hola", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	users := make(map[string]string)

	users["Axel"] = "alex@gmail.com"
	users["Julia"] = "Julia@gmail.com"

	if correo, ok := users["Julia"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("no fue posible obtener el valor")
	}

}
