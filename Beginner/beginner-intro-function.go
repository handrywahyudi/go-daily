package main

import (
	"fmt"
)

func average(number []float64) float64 {
	// panic("Not implemented")
	total := 0.0
	for _, v := range number {
		total += v
	}
	return total / float64(len(number))
}

func main() {
	number := []float64{5, 5, 5, 5, 5}
	fmt.Println(average(number))
}
