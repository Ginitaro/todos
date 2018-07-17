package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	v := r.Form

	fmt.Println(r)

	// Handle DB changes
	dbErr := model.RemoveTodoList("TodoBucket", 1)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Remove Successful!"))
	}

}
