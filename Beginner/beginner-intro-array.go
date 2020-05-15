package main

import (
	"fmt"
)

func main() {
	x := [5]float64{23, 40, 44, 923, 11}
	// var x [5]float64
	// x[0] = 23
	// x[1] = 40
	// x[2] = 33
	// x[3] = 80
	// x[4] = 90
	fmt.Println(x)

	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	var total2 float64 = 0
	for _, value := range x {
		total2 += value
	}
	fmt.Println(total2 / float64(len(x)))
}
