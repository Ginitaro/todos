package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"todos/controller"
	"todos/database"
)

// Core function
func main() {

	database.Connect()

	// Initialize Router
	r := mux.NewRouter()

	// Handle routes
	r.HandleFunc("/", controller.IndexHandler).Methods("GET")
	r.HandleFunc("/create", controller.TodoListHandler).Methods("GET", "POST")
	r.HandleFunc("/{id}/create", controller.TodoHandler).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}
