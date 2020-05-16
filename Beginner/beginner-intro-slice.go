package main

import (
	"fmt"
)

func main() {
	x := [5]float64{1, 2, 3, 4, 5}
	result := x[0:3]
	fmt.Println(result)
}
