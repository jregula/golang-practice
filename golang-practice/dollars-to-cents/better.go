package main

import (
	"fmt"
	"math"
)

func main() {
	// Declare a variable for the number of dollars
	dollars := 2.87

	// Convert the dollars to cents by multiplying by 100 and rounding to the nearest integer
	cents := int(math.Round(dollars * 100))

	// Initialize variables to keep track of the number of each coin
	numQuarters := 0
	numDimes := 0
	numNickels := 0
	numPennies := 0

	// Use a loop to continually subtract the largest possible coin value from the total number of cents
	for cents > 0 {
		if cents >= 25 {
			numQuarters++
			cents -= 25
		} else if cents >= 10 {
			numDimes++
			cents -= 10
		} else if cents >= 5 {
			numNickels++
			cents -= 5
		} else {
			numPennies++
			cents--
		}
	}

	// Print the result
	fmt.Printf("%v dollars is equal to %v cents (%v pennies, %v nickels, %v dimes, %v quarters)", dollars, int(dollars*100), numPennies, numNickels, numDimes, numQuarters)
}
