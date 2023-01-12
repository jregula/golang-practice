package main

import (
	"fmt"
	"os"
	"encoding/json"
	"encoding/csv"
	"strconv"
	

)

type Vegetables struct {
	Vegetable string `json:"vegetable"`
	Fruit string `json:"fruit"`
	Rank int `json:"rank`
}

func main() {

	var vegetables []Vegetables

	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}
	
	if err := json.NewDecoder(jsonFile).Decode(&vegetables); err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	csvFile, err := os.Create("data.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	header := []string{"vegetable", "fruit", "rank"}

	csvwriter.Write(header)


	for _, vegetable := range vegetables {
		csvwriter.Write([]string{vegetable.Vegetable,vegetable.Fruit,strconv.Itoa(vegetable.Rank)})
	}

	defer csvwriter.Flush()




}