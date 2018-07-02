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
	r.HandleFunc("/", controller.Homepage).Methods("GET")
	r.HandleFunc("/get_todolist", controller.GetTodoList).Methods("GET")
	r.HandleFunc("/create", controller.TodoListCreate).Methods("GET", "POST")
	r.HandleFunc("/{id}/create", controller.TodoCreate).Methods("GET", "POST")
	r.HandleFunc("/{todolist_id}/{todo_id}/remove", controller.TodoRemove).Methods("POST")
	http.ListenAndServe(":3000", r)
}
