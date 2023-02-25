package figuras

import "fmt"

type Figura interface {
	Area() float64
	Perimetro() float64
}

func AreaFigura(figura Figura) {
	fmt.Println("Area: ", figura.Area())
}
func PerimetroFigura(figura Figura) {
	fmt.Println("Perimetro: ", figura.Perimetro())
}
