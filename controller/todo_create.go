package controller

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"gitlab.com/mariusbeineris/todos/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
		todo := Todo{
			Name: r.FormValue("name"),
		}

		// Convert id<string> from request to id<int>
		tododata_id, err1 := strconv.Atoi(mux.Vars(r)["id"])
		if err1 != nil {
			log.Fatal(tododata_id)
		}

		// Handle DB changes
		database.DBCon.Update(func(tx *bolt.Tx) error {

			// Get TodoBucket instance
			b := tx.Bucket([]byte("TodoBucket"))

			// Get TodoData JSON by id from TodoBucket
			v := b.Get([]byte(itob(tododata_id)))

			var todo_data_item TodoData

			// Unserialize JSON to Tododata instance
			if err := json.Unmarshal(v, &todo_data_item); err != nil {
				panic(err)
			}

			// Find next available id for autoincrementing
			if len(todo_data_item.Todos) > 0 {
				todo.ID = todo_data_item.Todos[len(todo_data_item.Todos)-1].ID + 1
			} else {
				todo.ID = 0
			}

			// Append Todo into Todos of the TodoData instance
			todo_data_item.Todos = append(todo_data_item.Todos, todo)

			// Serialize TodoData instance
			buf, err := json.Marshal(todo_data_item)
			if err != nil {
				return err
			}

			// Update TodoData record in DB
			return b.Put(itob(todo_data_item.ID), buf)
		})

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}
