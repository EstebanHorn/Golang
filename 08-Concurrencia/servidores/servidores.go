package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		//fmt.Println(servidor, " no esta disponible")
		canal <- servidor + " no esta disponible"
	} else {
		canal <- servidor + " esta funcionando correctamente"
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"https://oregoom.com",
		"https://youtube.com",
		"https://udemy.com",
		"https://www.gooasdgle.com/",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	final := time.Since(inicio)
	fmt.Println("tiempo de ejecucion: ", final)
}
