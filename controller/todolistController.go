package controller

import (
	"encoding/json"
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
	dbErr, newTodoList := model.TodoListUpdate("TodoBucket", todo_data)
	if dbErr == nil {

		todolist, err := json.Marshal(newTodoList)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(todolist)
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

	// Convert form todolist_id to integer
	id, err := strconv.Atoi(r.Form.Get("todolist_id"))
	if err != nil {
		panic(err)
	}

	// Handle DB changes
	dbErr := model.RemoveTodoList("TodoBucket", id)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Remove Successful!"))
	}

}
