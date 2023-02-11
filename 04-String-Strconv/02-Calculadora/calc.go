package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(exp string) int {

	expArray := strings.Split(exp, "+")
	var result int

	for _, valor := range expArray {
		num, e := strconv.Atoi(valor)
		if e != nil {
			fmt.Println("ERROR: ", e)
			break
		} else {
			result += num
		}
	}

	return result
}

func main() {

	var exp string

	fmt.Println("Ingrese una suma: ")
	fmt.Scanln(&exp)

	fmt.Println("el resultado es: ", sumar(exp))

}
