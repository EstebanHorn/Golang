package figuras

type Cuadrado struct {
	Ancho  int
	Altura int
}

func (c *Cuadrado) Area() float64 {
	return (float64(c.Altura) * float64(c.Ancho))
}

func (c *Cuadrado) Perimetro() float64 {
	return (2*float64(c.Altura) + 2*float64(c.Ancho))
}
