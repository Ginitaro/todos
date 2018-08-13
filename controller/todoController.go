package controller

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"todos/errorhandler"
	"todos/model"
)

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// Create instance of Todo based on form input
	todo := model.Todo{
		Name: r.FormValue("name"),
	}

	// Convert id<string> from request to id<int>
	tododata_id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(tododata_id)
	}

	// Handle DB changes
	dbErr := model.TodoUpdate("TodoBucket", tododata_id, todo)
	if dbErr == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("200 - Great Success!"))
	}
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./view/todo_remove.html"))

	// Convert id<string> from request to id<int>
	todolist_id, err := strconv.Atoi(mux.Vars(r)["todolist_id"])
	if err != nil {
		log.Fatal(todolist_id)
	}

	todo_id, err := strconv.Atoi(mux.Vars(r)["todo_id"])
	if err != nil {
		log.Fatal(todo_id)
	}

	// Handle DB changes
	dbErr := model.TodoRemove("TodoBucket", todolist_id, todo_id)
	if dbErr != nil {
		errorhandler.CatchError(dbErr, "Todo not found.")
	} else {
		tmpl.Execute(w, struct{ Success bool }{true})
	}

}

func TodoToggle(w http.ResponseWriter, r *http.Request) {

	// Convert id<string> from request to id<int>
	todolist_id, err := strconv.Atoi(mux.Vars(r)["todolist_id"])
	if err != nil {
		log.Fatal(todolist_id)
	}

	todo_id, err := strconv.Atoi(mux.Vars(r)["todo_id"])
	if err != nil {
		log.Fatal(todo_id)
	}

	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		log.Fatal(status)
	}

	// Handle DB changes
	dbErr := model.TodoToggle("TodoBucket", todolist_id, todo_id, status)
	if dbErr != nil {
		errorhandler.CatchError(dbErr, "Todo not found.")
	}

}
