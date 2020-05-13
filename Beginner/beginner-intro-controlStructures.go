package main

import (
	"fmt"
)

func main() {

	// For loop
	for f := 1; f <= 100; f++ {
		fmt.Println(f)
	}
	fmt.Println("--------------------------------------")

	// IF Statement
	a := 2
	if a == 2 {
		fmt.Println("A is 2.")
	} else {
		fmt.Println("A is not 2.")
	}
	fmt.Println("--------------------------------------")

	// Implement For and if in one block
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Number of %v is even.\n", i)
		} else {
			fmt.Printf("Number of %v is odd.\n", i)
		}
	}
	fmt.Println("--------------------------------------")

	// Switch
	n := 4
	switch n {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Unknown number")
	}

	fmt.Println("--------------------------------------")

	// Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
	for q := 1; q <= 100; q++ {
		if q%3 == 0 {
			fmt.Println(q)
		}
	}

	// FizBuzz Program
	for b := 1; b <= 100; b++ {
		if b%3 == 0 && b%5 == 0 {
			fmt.Println("FizBuzz")
		} else if b%5 == 0 {
			fmt.Println("Buzz")
		} else if b%3 == 0 {
			fmt.Println("Fiz")
		} else {
			fmt.Println(b)
		}
	}
	fmt.Println("--------------------------------------")

}
