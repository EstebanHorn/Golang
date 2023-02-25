package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

type Figura interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	radio float64
}
type Cuadrado struct {
	ancho  int
	altura int
}

func (c *Circulo) Area() float64 {
	return (pi * (c.radio * c.radio))
}
func (c *Cuadrado) Area() float64 {
	return (float64(c.altura) * float64(c.ancho))
}

func (c *Circulo) Perimetro() float64 {
	return (2 * pi * c.radio)
}
func (c *Cuadrado) Perimetro() float64 {
	return (2*float64(c.altura) + 2*float64(c.ancho))
}

func areaFigura(figura Figura) {
	fmt.Println("Area: ", figura.Area())
}
func perimetroFigura(figura Figura) {
	fmt.Println("Perimetro: ", figura.Perimetro())
}

func main() {

	circulo1 := Circulo{
		radio: 3,
	}

	cuadrado1 := Cuadrado{
		ancho:  2,
		altura: 4,
	}

	perimetroFigura(&circulo1)
	areaFigura(&circulo1)

	perimetroFigura(&cuadrado1)
	areaFigura(&cuadrado1)

}
