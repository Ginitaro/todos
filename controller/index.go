package controller

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"html/template"
	"net/http"
	"todos/database"
)

type Todo struct {
	ID   int
	Name string
	Done bool
}

type TodoData struct {
	ID    int
	Title string
	Todos []Todo
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./view/layout.html"))
	var todo_data []TodoData

	database.DBCon.View(func(tx *bolt.Tx) error {

		// Get TodoBucket instance
		b := tx.Bucket([]byte("TodoBucket"))
		c := b.Cursor()

		// Get all data in TodoBucket
		for k, v := c.First(); k != nil; k, v = c.Next() {

			var todo_data_item TodoData

			// Unseriale each row in TodoBucket
			if err := json.Unmarshal(v, &todo_data_item); err != nil {
				panic(err)
			}

			// Append unserialized TodoData into array
			todo_data = append(todo_data, todo_data_item)
		}

		return nil
	})

	tmpl.Execute(w, todo_data)
}
