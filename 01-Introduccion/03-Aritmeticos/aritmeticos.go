package main

import "fmt"

func main() {
	a := 10
	b := 20

	//suma
	result := a + b
	fmt.Println(result)
	//resta
	result = a - b
	fmt.Println(result)
	//producto
	result = a * b
	fmt.Println(result)
	//division
	result = b / a
	fmt.Println(result)
	//division exacta
	var div float64 = 2.0 / 3.0
	fmt.Println(div)
	//modulo
	result = a % b
	fmt.Println(result)
}
