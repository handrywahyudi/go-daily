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

// Use Variadic Function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Use closure function
func makeEvenNumber() func() uint {
	i := uint(0)
	return func() (num uint) {
		num = i
		i += 2
		return
	}
}

// Use recursion function
func factorial(z uint) uint {
	if z == 0 {
		return 1
	}
	return z * factorial(z-1)
}

// Use defer
func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func main() {
	number := []float64{5, 5, 5, 5, 5}
	fmt.Println(average(number))

	fmt.Println(add(add(1, 2, 3)))

	// Use closure
	addTwoNumbers := func(x, y int) int {
		return x + y
	}
	fmt.Println(addTwoNumbers(3, 3))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenNumber()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	// Use recursion
	fmt.Println(factorial(2))

	// Use defer
	// defer second()
	// first()

	// Use panic and recover
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
