package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"todo/todo"
)

const (
	port = "8000"
	host = "127.0.0.1"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := todo.NewData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			d := json.NewDecoder(r.Body)
			t := &todo.TodoAdd{}
			if err := d.Decode(t); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data.AddTodo(t.Title)
			w.Write([]byte("Todo added"))
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			d := json.NewDecoder(r.Body)
			t := &todo.TodoUpdate{}
			if err := d.Decode(t); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data.UpdateTodo(t.Id)
			w.Write([]byte("Todo changed"))
		}
	})

	addr := fmt.Sprintf("%s:%s", host, port)

	log.Println(http.ListenAndServe(addr, nil))
}
