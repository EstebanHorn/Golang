package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Usuario struct {
	nombre string
	edad   int
}

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func renderTemplate(res http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(res, name, data)

	if err != nil {
		//http.Error(res, "No es posible retornar el template", http.StatusInternalServerError)
		HError(res, http.StatusInternalServerError)
	}
}

func HError(res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	errorTemplate.Execute(res, nil)

}
func Index(res http.ResponseWriter, req *http.Request) {

	//template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	usuario := Usuario{"Esteban", 19}
	renderTemplate(res, "index.html", usuario)

}
func Registro(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "registro.html", nil)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("El servidor esta corriendo en puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
