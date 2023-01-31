// https://www.youtube.com/watch?v=YMFkgN9r_jg

// Focus on learning http library

package main
import (
    // "fmt"
    "net/http"
	"log"
	"encoding/json"
)

type Book struct {
	Title string `json:"title"`
	Page int `json:"page"`
}

func test(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book {
		Title: "Test Title",
		Page: 111,
	}
	json.NewEncoder(w).Encode(book)
}

func main() {
	http.HandleFunc("/hello", test)
	log.Fatal(http.ListenAndServe(":5100", nil))
}