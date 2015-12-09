package main

import (
	"os"

	"github.com/wawandco/gontact/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/wawandco/gontact/Godeps/_workspace/src/github.com/gorilla/mux"
	"github.com/wawandco/gontact/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contact", handlers.ContactHandler).Methods("POST")

	//TODO: Middlewares add validation logic, could be with JWT or token authentication ?
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + serverPort())
}

func serverPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
