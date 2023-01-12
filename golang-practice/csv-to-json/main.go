package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"encoding/json"
)

type Vegetable struct {
	Vegetable string
	Fruit string
	Rank string
}
var vegetables []*Vegetable


func CSVtoJSON(source, destination string) error {
	fmt.Printf("translating from %s to %s\n", source, destination)

	// open CSV File
	csvFile, err := os.Open(source)

	// Return error if file doesn't exist
	if err != nil {
		return err
	}
	// Close at the end of function
	defer csvFile.Close()

	// 
	csvContent := csv.NewReader(csvFile)

	// skips first line
    if _, err := csvContent.Read(); err != nil {
        return  err
    }

	content, err := csvContent.ReadAll()

	if err != nil {
		return err
	}

	for _, i := range content {
		vegetables = append(vegetables, &Vegetable{Vegetable:i[0], Fruit: i[1], Rank: i[2] }) 
	}
	p, _ := json.Marshal(vegetables)

	outputFile, outputError := os.OpenFile("data.json", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
	  fmt.Printf("An error occurred with file creation\n")
	  return outputError
	}

	defer outputFile.Close()

	err = os.WriteFile("data.json", p, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}





func main() {
	if err := CSVtoJSON("data.csv", "data.json"); err != nil {
		log.Fatal(err)
	}
}