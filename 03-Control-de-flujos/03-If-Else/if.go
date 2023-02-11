package main

import "fmt"

func main() {

	var consumo, descuento float64
	//var datosDescuento string

	fmt.Println("Ingrese consumo")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			descuento = consumo * 0.1
		} else if consumo <= 200 {
			descuento = consumo * 0.2
		} else {
			descuento = consumo * 0.3
		}

		iva := 1.21
		fmt.Println("El descuento es de:", descuento)
		totalAPagar := (consumo - descuento) * iva

		fmt.Println("El total con iva es:", totalAPagar)
		fmt.Println("El descuento es de:", descuento)

	} else {
		fmt.Println("El consumo deebe ser mayo a cero")
	}
}
