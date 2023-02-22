package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func main() {
	persona1 := Persona{"Esteban", 20}

	fmt.Println(persona1)

	persona1.nombre = "Samuel"
	fmt.Println(persona1)

	persona2 := Persona{
		nombre: "Julia",
		edad:   27,
	}
	fmt.Println(persona2)

}
