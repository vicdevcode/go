package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Println(err)
		}
	})

	log.Println(http.ListenAndServe(":5000", nil))
}
