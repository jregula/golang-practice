package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/{input}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        input := vars["input"]
        fmt.Fprintf(w, "Input received: %s", input)
    }).Methods("GET")
    http.ListenAndServe(":8000", r)
}