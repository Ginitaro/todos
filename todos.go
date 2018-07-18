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

	//r.HandleFunc("/", controller.Homepage).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/get_todolist", controller.GetTodoList).Methods("GET")
	api.HandleFunc("/create", controller.TodoListCreate).Methods("POST")
	api.HandleFunc("/remove", controller.RemoveTodoList).Methods("POST")
	api.HandleFunc("/{id}/create", controller.TodoCreate).Methods("POST")
	api.HandleFunc("/{todolist_id}/{todo_id}/remove", controller.TodoRemove).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("view/dist/")))
	r.PathPrefix("/").HandlerFunc(controller.GetTodoList)

	http.ListenAndServe(":3000", r)
}
