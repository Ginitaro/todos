package controller

import (
	"encoding/binary"
	"gitlab.com/mariusbeineris/todos/database"
	"html/template"
	"net/http"
)

func TodoDataHandler(w http.ResponseWriter, r *http.Request) {

	// Load template
	tmpl := template.Must(template.ParseFiles("tododata_create.html"))

	// Check if request is POST
	if r.Method != http.MethodPost {

		// Render view
		tmpl.Execute(w, nil)
		return
	} else {

		// Create instance of TodoData based on form input
		todo_data := TodoData{
			Title: r.FormValue("title"),
		}

		// Handle DB changes
		err := database.UpdateTodoData("TodoBucket", todo_data)

		if err != nil {
			tmpl.Execute(w, struct{ Success bool }{true})
		}
	}
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
