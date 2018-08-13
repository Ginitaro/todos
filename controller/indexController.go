package controller

import (
	// "encoding/json"
	// "github.com/boltdb/bolt"
	"html/template"
	"net/http"
	// "todos/database"
	// "todos/model"
)

func Homepage(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.New("index.html").Delims("#$%", "%$#").ParseFiles("./view/dist/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.Execute(w, nil)
}
