package handlers

import (
	"apiRest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(res)
	} else {
		models.SendData(res, users)
	}
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		models.SendData(res, user)
	}
}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	//obtener registro
	user := models.User{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(res)
	} else {
		user.Save()
		models.SendData(res, user)
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var userId int64

	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(res)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(res, user)
	}
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	if user, err := getUserByRequest(req); err != nil {
		models.SendNotFound(res)
	} else {
		user.Delete()
		models.SendData(res, user)
	}
}

func getUserByRequest(req *http.Request) (models.User, error) {
	vars := mux.Vars(req)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
