package main

import "fmt"

// Uno a uno
type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

// ----------------------------------------

// uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

// ------------------------------------------

func main() {
	/* uno a uno
	user1 := User{
		nombre: "Esteban",
		email:  "Esteban@gmail.com",
	}

	user2 := User{
		nombre: "Samuel",
		email:  "Samuel@gmail.com",
		activo: true,
	}

	user1Student := Student{
		user:   user1,
		codigo: "3n23k",
	}
	fmt.Println(user1)

	fmt.Println(user2)

	fmt.Println(user1Student)*/

	video1 := Video{titulo: "Introduccion"}
	video2 := Video{titulo: "Instalancion"}

	curso1 := Curso{
		titulo: "Curso GOlang",
		videos: []Video{video1, video2},
	}

	video1.curso = curso1
	video2.curso = curso1

	for _, video := range curso1.videos {
		fmt.Println(video.titulo)
	}

}
