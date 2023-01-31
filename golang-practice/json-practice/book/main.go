package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func main() {
	var book Book

	jsonFile, err := os.Open("book.json")
	if err != nil {
		fmt.Println(err)
	}

	d := json.NewDecoder(jsonFile)

	if err := d.Decode(&book); err != nil {
		fmt.Println(err)
	}

	fmt.Println(book)
}

type Book struct {
	Book []Books `json:"book"`
}


type Books struct {
	Id string `json:"id"`
	Language string `json:"language"`
	Edition string `json:"edition"`
	Author string `json:"author"`

}