package main

import (
	"fmt"
)

func unamedMinMax(x, y int) (int, int) {
	if x > y {
		max := x
		min := y
		return max, min
	} else {
		max := y
		min := x
		return max, min
	}
}

func minMax(x, y int) (min, max int) {
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	return min, max
}

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		max = x
		min = y
	} else {
		max = y
		min = x
	}
	return
}

func sort(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}

func main() {
	// anonymous function
	y := 4
	square := func(s int) int {
		return s * s
	}
	fmt.Println("Square of ", y, " is ", square(y))

	square = func(s int) int {
		return s + s
	}
	fmt.Println("Square of ", y, " is ", square(y))

	fmt.Println(minMax(5, 6))
	fmt.Println(namedMinMax(5, 6))
	min, max := namedMinMax(23, -6)
	fmt.Println(min, max)
}
