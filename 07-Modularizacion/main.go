package main

import (
	"fmt"
	"paquetes/models"
)

func main() {
	/*
		cua1 := figuras.Cuadrado{
			Ancho:  2,
			Altura: 4,
		}

		cir1 := figuras.Circulo{
			Radio: 2,
		}

		figuras.AreaFigura(&cua1)
		figuras.AreaFigura(&cir1)
		figuras.PerimetroFigura(&cua1)
		figuras.PerimetroFigura(&cir1)*/

	p1 := models.Persona{}
	p1.Constructor("Esteban", 19)

	fmt.Println(p1)

}
