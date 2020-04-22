package main

import (
	"fmt"
	"os"
	"strconv"
)

// Write a Go program that keeps reading integers until you give the number 0 as input,
// then it prints the minimum and maximum integer in the input.
func main() {
	arguments := os.Args
	numbers := []int{}

	for i := 1; i < len(arguments); i++ {
		x, _ := strconv.Atoi(arguments[i])
		numbers = append(numbers, x)

	}
	fmt.Println(numbers)

	var n, biggest, smallest int
	for _, number := range numbers {
		if number > n {
			n = number
			biggest = n
		}
	}

	for _, number := range numbers {
		if number < n {
			n = number
			smallest = number
		}
	}

	fmt.Printf("The oldest in this organization is : %v years old.\n", biggest)
	fmt.Printf("The youngest in this organization is : %v years old.\n", smallest)
}
