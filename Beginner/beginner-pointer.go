package main

import "fmt"

func withPointer(x *int) {
	*x = *x * *x
}

type complex struct {
	x, y int
}

func newComplex(x, y int) *complex {
	return &complex{x, y}
}

func main() {
	x := -2
	withPointer(&x)
	fmt.Println(x)

	w := newComplex(4, -5)
	fmt.Println(*w)
	fmt.Println(w)

	// Another example of pointer
	var p *int

	if p != nil {
		fmt.Println("Value of p: ", *p)
	} else {
		fmt.Println("P is nil")
	}

	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("Value of pi: ", *p)
	} else {
		fmt.Println("P is nil")
	}

	var value float64 = 43.23
	fmt.Println("Value is ", value)
	pointer := &value
	*pointer = 34.3
	fmt.Println("Pointer is ", *pointer)

	*pointer = *pointer / 30
	fmt.Println("Pointer is ", *pointer)

}
