package main

import (
	"fmt"
	"log"
	"net/http"
)

type numberDumper int

func (n numberDumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here's your number: %d\n", n)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	h := http.NewServeMux()

	h.Handle("/one", numberDumper(1))
	h.Handle("/two", numberDumper(2))
	h.Handle("/three", numberDumper(3))
	h.Handle("/four", numberDumper(4))
	h.Handle("/five", numberDumper(5))

	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		fmt.Fprintln(w, "That's not a supported number!")
	})

	hl := logger(h)

	err := http.ListenAndServe(":9999", hl)
	log.Fatal(err)
}
