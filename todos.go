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

	r.HandleFunc("/", controller.Homepage).Methods("GET")

	// Handle routes
	r.HandleFunc("/api/get_todolist", controller.GetTodoList).Methods("GET")
	r.HandleFunc("/api/create", controller.TodoListCreate).Methods("POST")
	r.HandleFunc("/api/remove", controller.RemoveTodoList).Methods("POST")
	r.HandleFunc("/api/{id}/create", controller.TodoCreate).Methods("POST")
	r.HandleFunc("/api/{todolist_id}/{todo_id}/remove", controller.TodoRemove).Methods("POST")

	http.ListenAndServe(":3000", r)
}
