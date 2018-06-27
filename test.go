package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Struct variables needs to be captitalized to be public
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

var db *bolt.DB

func index_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	var todo_data []TodoData

	db.View(func(tx *bolt.Tx) error {

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

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func tododata_create_handler(w http.ResponseWriter, r *http.Request) {

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
		db.Update(func(tx *bolt.Tx) error {

			// Get TodoBucket instance
			b := tx.Bucket([]byte("TodoBucket"))

			// Find next available id for autoincrementing
			id, _ := b.NextSequence()
			todo_data.ID = int(id)

			// Serialize TodoData instance
			buf, err := json.Marshal(todo_data)
			if err != nil {
				return err
			}

			// Update TodoData record in DB
			return b.Put(itob(todo_data.ID), buf)
		})

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func todo_create_handler(w http.ResponseWriter, r *http.Request) {

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
		tododata_id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			log.Fatal(tododata_id)
		}

		// Handle DB changes
		db.Update(func(tx *bolt.Tx) error {

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
			id, _ := b.NextSequence()
			todo.ID = int(id)

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

// Core function
func main() {

	// Open the DB connection
	var err error
	db, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize TodoBucket bucket
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("TodoBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_ = b
		return nil
	})

	// Dev, print all data from DB in TodoBucket
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodoBucket"))
		v := b.Get([]byte(itob(1)))
		fmt.Printf("The title is: %s\n", v)
		return nil
	})

	// Initialize Router
	r := mux.NewRouter()
	r.HandleFunc("/", index_handler).Methods("GET")
	r.HandleFunc("/create", tododata_create_handler).Methods("GET", "POST")
	r.HandleFunc("/{id}/create", todo_create_handler).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}
