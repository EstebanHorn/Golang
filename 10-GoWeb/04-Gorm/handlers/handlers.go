package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(res, users, http.StatusOK)
}

func GetUser(res http.ResponseWriter, req *http.Request) {
	if user, err := getUserById(req); err != nil {
		sendError(res, http.StatusNotFound)
	} else {
		sendData(res, user, http.StatusOK)
	}

}

func CreateUser(res http.ResponseWriter, req *http.Request) {

	//obtener registro
	user := models.User{}
	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(res, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(user)
		sendData(res, user, http.StatusCreated)
	}
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	var userId int64

	if user_ant, err := getUserById(req); err != nil {
		sendError(res, http.StatusNotFound)
	} else {
		userId = user_ant.Id

		user := models.User{}
		decoder := json.NewDecoder(req.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(res, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(user)
			sendData(res, user, http.StatusOK)
		}
	}

}

func DeleteUser(res http.ResponseWriter, req *http.Request) {

	if user, err := getUserById(req); err != nil {
		sendError(res, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(res, user, http.StatusOK)
	}
}

func getUserById(req *http.Request) (models.User, *gorm.DB) {
	user := models.User{}
	vars := mux.Vars(req)
	userId, _ := strconv.Atoi(vars["id"])
	if err := db.Database.First(user, &userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}

}
