package main

import (
	"fmt"
    "strconv"

)

func main() {
	// user input
    var first string
    fmt.Println("Enter a binary string: ")
    fmt.Scanln(&first)
	decimal, err := strconv.ParseInt(first,2,64)

	if len(first) > 8 {
		fmt.Println("Max len 8")
	} else if err != nil {
		fmt.Println("Please only enter 1 or 0")
	} else {
		fmt.Println(decimal)
	}
	

}