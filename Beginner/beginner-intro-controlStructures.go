package main

import (
	"fmt"
)

func main() {

	// For loop
	for f := 1; f <= 100; f++ {
		fmt.Println(f)
	}

	// IF Statement
	a := 2
	if a == 2 {
		fmt.Println("A is 2.")
	} else {
		fmt.Println("A is not 2.")
	}

	// Implement For and if in one block
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Number of %v is even.\n", i)
		} else {
			fmt.Printf("Number of %v is odd.\n", i)
		}
	}

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
	}

}
