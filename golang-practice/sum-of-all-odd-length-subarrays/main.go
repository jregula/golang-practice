package main

import "fmt"

func sumOddLengthSubarrays(A []int) int {
    sum := 0
    for i := 0; i < len(A); i++ {
        for j := i; j < len(A); j += 2 {
			fmt.Println(A[i:j+1])
            sum += sumSlice(A[i:j+1])
        }
    }
    return sum
}

func sumSlice(slice []int) int {
    sum := 0
    for _, v := range slice {
        sum += v
    }
    return sum
}

func main() {
	testArray := []int{1,51,1,2,3,45}
	sumOddLengthSubarrays(testArray)
}