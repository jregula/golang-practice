package main

// https://github.com/florinpop17/app-ideas/blob/master/Projects/1-Beginner/Dollars-To-Cents-App.md
// convert user input of dollars to cents

import (
	"fmt"
	"log"
    "math"

)


func convertDollartoCent() error {
	var test float64

	fmt.Println("Enter a float64: ")
	fmt.Scanln(&test)
	coins := map[float64]int{
		0.2: 0,
		0.1: 0,
		0.01: 0,
	}
	integer, frac := math.Modf(test)
	frac = math.Round(frac*100)

	coins[0.2] = int(integer / 0.2)

	x :=1
	for x == 1 {
		if math.Mod(frac, 20) == 0 {
			coins[0.2] = coins[0.2] + int(math.Mod(frac, 20))
			x = 0
		} else {
			coins[0.2] = coins[0.2] + int(frac / 20)
			frac = math.Mod(frac, 20)
			// fmt.Println(frac)
			// 4
		}
		if math.Mod(frac, 10) == 0 {
			coins[0.1] = int(math.Mod(frac, 20))
			x = 0
		} else if frac / 10 > 0 {
			coins[0.1] = int(frac / 10)
			frac = math.Mod(frac, 10)
			}
		coins[0.01] = int(frac / 1)
		x = 0

	}
	fmt.Println(coins)
	return nil
}

func main() {
	if err := convertDollartoCent(); err != nil {
		log.Fatal(err)
	}
}