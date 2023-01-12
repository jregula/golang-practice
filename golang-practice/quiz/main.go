package main

import (
	"quiz/parser"
    "flag"
	"time"
	"fmt"

)

func main() {
	currentChannel := make(chan parser.FuncResult, 1)
	

	csvFilepath := flag.String("csvfile", "problems.csv", "csv file with questions and answers")
	flag.Parse()

	go func() {
		txt :=parser.QuestionRange(*csvFilepath)

		currentChannel <- txt
	 }()


	select {
		case res := <-currentChannel:
			fmt.Printf("You answered %v answers correctly and %v incorrectly\n", res.Correct, res.Incorrect)
		case <-time.After(3 * time.Second):
			fmt.Println(res)

	}
	fmt.Println("Main function exited!")
}