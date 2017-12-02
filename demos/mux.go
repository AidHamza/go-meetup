package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit the blog page!")
	})

	h.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, you hit contact us!")
	})

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "You're lost, go home")
	})

	err := http.ListenAndServe(":8080", h)
	log.Fatal(err)
}