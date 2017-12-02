package main

import (
    "fmt"
    "net/http"
)

func helloGophers(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello Gophers!")
}

func main() {
    http.HandleFunc("/", helloGophers)
    http.ListenAndServe(":8080", nil)
}