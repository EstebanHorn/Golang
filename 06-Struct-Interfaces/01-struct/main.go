package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) imprimir() {
	fmt.Printf("\nNombre %s \nEdad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nombre string) {
	p.nombre = nombre
}

type Empleado struct {
	Persona
}

func main() {
	persona1 := Persona{"Esteban", 20}

	persona1.imprimir()
	persona1.editarNombre("Juan")
	persona1.imprimir()

	persona2 := Persona{
		nombre: "Julia",
		edad:   27,
	}

	persona2.imprimir()

	empleado1 := Empleado{}
	empleado1.nombre = "Jorge"
	empleado1.edad = 55
	empleado1.imprimir()

}
