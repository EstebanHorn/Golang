package main

import "fmt"

func main() {

	a := 2
	b := 3

	fmt.Println(!true)

	fmt.Println(true && true)

	fmt.Println(false || false)

	r := a == 2 && !(b == 2)

	fmt.Println(r)
}
