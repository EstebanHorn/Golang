package main

import (
	"fmt"
	"strings"
)

func reapeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	/*
		func() {
			fmt.Println("Anonimo")
		}()
		myfunc := func(nombre string) string {
			return fmt.Sprintf("Hola %s anonimo", nombre)
		}
		fmt.Println(myfunc("Esteban"))
	*/

	repeat3 := reapeat(3)
	fmt.Println(repeat3("Esteban \n"))

}
