package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World from Go!")
}

func main() {
    http.HandleFunc("/hello-go", helloWorld)
    http.ListenAndServe(":8080", nil)
}