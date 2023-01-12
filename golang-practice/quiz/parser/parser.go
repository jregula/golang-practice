package parser

import (
	"fmt"
	"os"
	"encoding/csv"
	"log"
)
// function to open file and save output to variable
func Reader(csvFilepath string) (contents [][]string, er error) {

	var fileContent [][]string
	file, err := os.Open(csvFilepath) // For read access.

	if err != nil {
		log.Fatal(err)
		return fileContent, err
	}
	defer file.Close()
	r := csv.NewReader(file)

	content, err := r.ReadAll()

	if err != nil {
		return fileContent, err
	}

	fileContent = content
	return fileContent, nil

}

type FuncResult struct {
	Err    error
	Correct int
	Incorrect int
}

// function to range through questions
func QuestionRange(csvFilepath string) FuncResult {
	var userAnswer string
	correct := 0
	content, err := Reader(csvFilepath)

	if err != nil {
		log.Fatal(err)
		return FuncResult{Err: err, Correct: 0, Incorrect: 0}
	}

	for _, row := range content {
		fmt.Printf("What is %v?\n",row[0])
		fmt.Scanln(&userAnswer)
		if row[1] == userAnswer {
			correct++
		}
	}
	return FuncResult{Err: nil, Correct: correct, Incorrect: len(content) - correct}

}