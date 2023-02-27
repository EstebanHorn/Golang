package main

import (
	"fmt"
	"goMySql/db"
	"goMySql/models"
)

func main() {
	db.ConnectDB()
	fmt.Println("usuarios:", models.ListUsers())
	//db.ResetTable("users")
	user := models.CreateUser("Julia", "111111", "Julia@gmail.com")
	user.Save()
	fmt.Println("usuarios:", models.ListUsers())

	//fmt.Println(user)S
	//db.CreateTable(models.UserSchema, "users")
	db.DiscconectDb()
}
