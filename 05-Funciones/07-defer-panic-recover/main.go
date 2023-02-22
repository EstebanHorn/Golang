package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if error := recover(); error != nil {
			fmt.Println("UPS! ah ocurrido un error")
		}
	}()

	if file, error := os.Open("hola.txt"); error != nil {
		panic("ERROR \n no fue posible abrir el archivo")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo...")
			file.Close()
		}()
		cont := make([]byte, 254)
		long, _ := file.Read(cont)
		contArch := string(cont[:long])
		fmt.Println(contArch)
	}

}
