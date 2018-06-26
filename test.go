package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// Struct variables needs to be captitalized to be public
type Todo struct {
	Title string
	Done  bool
}

type TodoData struct {
	Task  string
	Todos []Todo
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	todo_data := TodoData{
		Task: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, todo_data)
}

func create_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("create.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	todo_title := Todo{
		Title: r.FormValue("title"),
	}

	// do something with details
	_ = todo_title

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("TodoBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_ = b
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodoBucket"))
		err := b.Put([]byte("Title"), []byte("Test"))
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodoBucket"))
		v := b.Get([]byte("Title"))
		fmt.Printf("The title is: %s\n", v)
		return nil
	})

	r := mux.NewRouter()
	r.HandleFunc("/", index_handler).Methods("GET")
	r.HandleFunc("/create", create_handler).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}
