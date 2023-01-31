package main

import (
	"quiz/parser"
    "flag"
	// "time"
	// "fmt"

)

func main() {
	csvFilepath := flag.String("csvfile", "problems.csv", "csv file with questions and answers")
	flag.Parse()
	parser.QuestionRange(*csvFilepath)
}