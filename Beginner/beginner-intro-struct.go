package main

import (
	"fmt"
	"math"
)

// Struct
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// Method
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
