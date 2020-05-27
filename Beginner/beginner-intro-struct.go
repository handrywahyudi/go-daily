package main

import (
	"fmt"
)

type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.y, c.r)
}
