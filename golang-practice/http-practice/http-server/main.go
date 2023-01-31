package main

import (
	"fmt"
	"net/http"
	"http-server/api"
)

func main() {
	fmt.Println("Hello World!")

	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
