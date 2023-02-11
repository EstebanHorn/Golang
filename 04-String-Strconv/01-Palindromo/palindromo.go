package main

import (
	"fmt"
	"strings"
)

func reverse(frase string) string {

	Arrayfrase := strings.Split(frase, "")
	ArrayRev := make([]string, 0)

	for i := len(Arrayfrase) - 1; i >= 0; i-- {
		ArrayRev = append(ArrayRev, Arrayfrase[i])
	}

	return strings.Join(ArrayRev, "")

}

func palindromo(frase string) bool {

	frase = strings.ToLower(frase)
	frase = strings.ReplaceAll(frase, " ", "")
	fraseRev := reverse(frase)

	return (frase == fraseRev)

}

func main() {

	frase := "Añora la roña"

	if palindromo(frase) {
		fmt.Println(frase, "Es Palindromo")
	} else {
		fmt.Println(frase, "No es Palindromo")
	}

}
