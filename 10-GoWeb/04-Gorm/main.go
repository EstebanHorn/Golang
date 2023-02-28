package main

import (
	"gorm/models"
)

func main() {

	models.MigrarUsers()
	/*
		//Rutas
		mux := mux.NewRouter()

		//EndPoints
		mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
		mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
		mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

		//servidor
		log.Fatal(http.ListenAndServe(":3000", mux))*/
}
