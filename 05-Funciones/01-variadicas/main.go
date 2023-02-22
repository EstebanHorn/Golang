package main

import "fmt"

func sumar(nombre string, n ...int) (string, int) {
	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range n {
		total += num
	}
	return mensaje, total
}

func main() {
	men, res := sumar("Esteban", 3, 5)
	println(men, res)

}
