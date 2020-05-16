package main

import (
	"fmt"
)

func main() {
	// Append function
	x := []float64{1, 2, 3, 4, 5}
	y := append(x, 6, 7, 8)
	fmt.Println(y)

	// Copy function
	m := []int{1, 2, 3}
	n := make([]int, 2)
	copy(n, m)
	fmt.Println(m, n)
}
