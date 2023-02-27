package models

import (
	"goMySql/db"
)

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `CREATE TABLE users(
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(64) NOT NULL,
	email VARCHAR(50),
	create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// Construir Usuario
func NewUser(username, password, email string) *User {
	user := &User{UserName: username, Password: password, Email: email}
	return user
}

// Crear registro e insertar
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Insertar Registro
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	res, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	user.Id, _ = res.LastInsertId()
}

// Listar usuarios
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}
	return users
}

// Obtener un registro
func GetUser(Id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, Id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}
	return user
}

// Actualizar resgitro
func (user *User) updateUser() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)
}

// Guardar o editar regsitro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.updateUser()
	}
}

// Eliminar registro
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}
