package controller

import (
	"html/template"
	"net/http"
	"todos/model"
)

func TodoListHandler(w http.ResponseWriter, r *http.Request) {

	// Load template
	tmpl := template.Must(template.ParseFiles("todolist_create.html"))

	// Check if request is POST
	if r.Method != http.MethodPost {

		// Render view
		tmpl.Execute(w, nil)
		return
	} else {

		// Create instance of TodoList based on form input
		todo_data := model.TodoList{
			Title: r.FormValue("title"),
		}

		// Handle DB changes
		dbErr := model.UpdateTodoList("TodoBucket", todo_data)
		if dbErr == nil {
			tmpl.Execute(w, struct{ Success bool }{true})
		}
	}
}
