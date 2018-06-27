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
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TodoBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			var todo_data_item TodoData

			if err := json.Unmarshal(v, &todo_data_item); err != nil {
				panic(err)
			}

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
	tmpl := template.Must(template.ParseFiles("tododata_create.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	} else {

		todo_data := TodoData{
			Title: r.FormValue("title"),
		}

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TodoBucket"))

			id, _ := b.NextSequence()
			todo_data.ID = int(id)

			buf, err := json.Marshal(todo_data)
			if err != nil {
				return err
			}

			return b.Put(itob(todo_data.ID), buf)
		})

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func todo_create_handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("todo_create.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	} else {

		todo := Todo{
			Name: r.FormValue("name"),
		}

		// tododata_id := mux.Vars(r)["id"]

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TodoBucket"))

			v := b.Get([]byte(itob(1)))

			var todo_data_item TodoData

			if err := json.Unmarshal(v, &todo_data_item); err != nil {
				panic(err)
			}

			id, _ := b.NextSequence()
			todo.ID = int(id)

			todo_data_item.Todos = append(todo_data_item.Todos, todo)

			buf, err := json.Marshal(todo_data_item)
			if err != nil {
				return err
			}

			return b.Put(itob(todo_data_item.ID), buf)
		})

		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func main() {
	var err error
	db, err = bolt.Open("my.db", 0600, nil)
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

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("TodoBucket"))
		v := b.Get([]byte(itob(1)))
		fmt.Printf("The title is: %s\n", v)
		return nil
	})

	r := mux.NewRouter()
	r.HandleFunc("/", index_handler).Methods("GET")
	r.HandleFunc("/create", tododata_create_handler).Methods("GET", "POST")
	r.HandleFunc("/{id}/create", todo_create_handler).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}
