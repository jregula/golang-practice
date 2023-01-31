package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct {
	Id   int
    Name string
}

var user User


func example(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(&user)
	if err!= nil {
		fmt.Println(err)

	}
	writer.WriteHeader(http.StatusOK)

	// vars := mux.Vars(request)
	// input := vars["input"]
	// fmt.Fprintf(writer, "Input received: %s", "test")

}

func main() {
	user.Id = 1
	user.Name = "john"
	mux := http.NewServeMux()
	mux.HandleFunc("/", example)
    http.ListenAndServe(":8000", mux)

}