package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"gitlab.com/mariusbeineris/todos/controller"
	"gitlab.com/mariusbeineris/todos/database"
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

// Core function
func main() {

	// Open the DB connection
	var err error
	database.DBCon, err = bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer database.DBCon.Close()

	// Initialize TodoBucket bucket
	database.DBCon.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("TodoBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_ = b
		return nil
	})

	// Dev, print all data from DB in TodoBucket
	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("TodoBucket"))
	// 	v := b.Get([]byte(itob(4)))
	// 	fmt.Printf("The title is: %s\n", v)
	// 	return nil
	// })

	database.DBCon.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("TodoBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	// Initialize Router
	r := mux.NewRouter()
	r.HandleFunc("/", controller.IndexHandler).Methods("GET")
	r.HandleFunc("/create", controller.TodoDataHandler).Methods("GET", "POST")
	r.HandleFunc("/{id}/create", controller.TodoHandler).Methods("GET", "POST")
	http.ListenAndServe(":3000", r)
}
