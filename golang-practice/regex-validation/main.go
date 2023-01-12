// https://github.com/florinpop17/app-ideas/blob/master/Projects/1-Beginner/Word-Frequency-App.md
package main

import (
	"fmt"
	"log"
	"regexp"
)
// Length of at least 6
// The first would require the user to enter five (5) capital letters, six (6) symbols and two hyphens (-) in any order. 
// This could be used as a password.
// The second which could be used as username would require the user to enter letters without spaces
// The third which could be used as email address would require the user to enter only email addresses on gmail 

var str string = "Hello123!"

func calculateRegex(str string) error{
	r, err := regexp.Compile("[0-9a-zA-Z]{6,}")

	if err != nil {
		return err
	}
    fmt.Println(r.MatchString(str))

	return nil
}

func main() {

	if err := calculateRegex(str); err != nil {
		log.Fatal(err)
	}
}