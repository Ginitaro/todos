package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todos/errorhandler"
	"todos/model"
)

func TodoListCreate(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var todo_data model.TodoList
	err := decoder.Decode(&todo_data)
	if err != nil {
		panic(err)
	}

	// Handle DB changes
	dbErr := model.TodoListUpdate("TodoBucket", todo_data)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Great Success!"))
	}

}

func GetTodoList(w http.ResponseWriter, r *http.Request) {

	todolist, err := model.GetTodoList("TodoBucket")

	if err != nil {
		errorhandler.CatchError(err, "Could not fetch todos list.")
	}

	bs, err := json.Marshal(todolist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)

}

func RemoveTodoList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["todolist_id"]

	i, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.CatchError(err, "Could convert id to integer.")
	}
	// Handle DB changes
	dbErr := model.RemoveTodoList("TodoBucket", i)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Remove Successful!"))
	}

}
