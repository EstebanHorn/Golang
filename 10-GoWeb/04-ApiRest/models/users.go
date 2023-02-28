package models

import (
	"apiRest/db"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
func ListUsers() (Users, error) {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}

	if rows, err := db.Query(sql); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
			users = append(users, user)
		}
		return users, nil
	}
}

// Obtener un registro
func GetUser(Id int) (*User, error) {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	if rows, err := db.Query(sql, Id); err != nil {
		return nil, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		}
		return user, nil
	}

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
