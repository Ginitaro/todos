package controller

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"todos/database"
)

func TodoHandler(w http.ResponseWriter, r *http.Request) {

	// Load template
	tmpl := template.Must(template.ParseFiles("todo_create.html"))

	// Check if request is POST
	if r.Method != http.MethodPost {

		// Render view
		tmpl.Execute(w, nil)
		return
	} else {

		// Create instance of Todo based on form input
		todo := database.Todo{
			Name: r.FormValue("name"),
		}

		// Convert id<string> from request to id<int>
		tododata_id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err == nil {
			log.Fatal(tododata_id)
		}

		// Handle DB changes
		dbErr := database.UpdateTodo("TodoBucket", tododata_id, todo)
		if dbErr == nil {
			tmpl.Execute(w, struct{ Success bool }{true})
		}
	}
}
