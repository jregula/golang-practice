// https://github.com/florinpop17/app-ideas/blob/master/Projects/1-Beginner/Word-Frequency-App.md
package main

import (
	"fmt"
	"strings"
	"log"
)

var str string = "hello hello hello bye bye what"

func calculateString(str string) (map[string]int, error){
	stringMap := make(map[string]int)
	stringList := strings.Fields(str)

	for _, s := range stringList {
		stringMap[s]++
	}
	return stringMap, nil
}

func main() {

	if stringMap, err := calculateString(str); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(stringMap)
	}


}