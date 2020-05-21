package main

import (
	"fmt"
)

func main() {
	number := []float64{43, 23, 45, 63, 21}

	total := 0.0
	for _, v := range number {
		total += v
	}
	fmt.Println(total / float64(len(number)))
}
