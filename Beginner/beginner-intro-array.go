package main

import (
	"fmt"
)

func main() {
	var x [5]float64
	x[0] = 23
	x[1] = 40
	x[2] = 33
	x[3] = 80
	x[4] = 90
	fmt.Println(x)

	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}
