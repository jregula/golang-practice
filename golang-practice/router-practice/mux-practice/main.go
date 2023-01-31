// learn how to use json with requests and mux


package main

import (
	// "github.com/gorilla/mux"
	"net/http"
	"fmt"
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

		vars := mux.Vars(request)
        input := vars["input"]
        fmt.Fprintf(writer, "Input received: %s", input)

}

func main() {
	user.Id = 1
	user.Name = "john"
    router := mux.NewRouter()
	router.HandleFunc("/{input}", example).Methods("GET")
    http.ListenAndServe(":8000", router)

}