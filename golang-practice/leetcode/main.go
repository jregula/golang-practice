package main

import (
    "errors"
    "fmt"
)

func testError(f int) (int, error) {
	if f > 5 {
		return -1, errors.New("Error")
	}
	return f, nil
}

func main() {

	if _, err := testError(6); err != nil {
		fmt.Println("err")
	}




}