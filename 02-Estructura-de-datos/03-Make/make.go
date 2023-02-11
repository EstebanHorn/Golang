package main

import "fmt"

func main() {

	num := make([]int, 3, 3)

	fmt.Println(num, len(num), cap(num))

	num[0] = 1
	num[1] = 2
	num[2] = 3

	num = append(num, 4)

	fmt.Println(num, len(num), cap(num))

}
