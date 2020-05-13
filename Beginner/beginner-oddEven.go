package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Number of %v is even.\n", i)
		} else {
			fmt.Printf("Number of %v is odd.\n", i)
		}
	}
}
