package main

import (
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/wawandco/gontact/handlers"
	"github.com/wawandco/gontact/middlewares"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contact", handlers.ContactHandler).Methods("POST")

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(middlewares.CORSMiddleware))
	n.Use(negroni.HandlerFunc(middlewares.TokenMiddleware))
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
