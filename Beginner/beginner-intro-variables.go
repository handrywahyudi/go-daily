package main

import (
	"fmt"
)

var dogsName string = "dog"

func dog() {
	fmt.Println("i have a", dogsName)
}

func cToFah(c float64) float64 {
	total := (c * 9.0 / 5.0) + 32.0
	return total
}

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	y := "first"
	fmt.Println(y)
	y += " second"
	fmt.Println(y)

	fmt.Println("We have a", dogsName)
	dog()

	// const name string = "handry"
	// name = "wahyudi"
	// fmt.Println(name)

	/*
	   Create program to convert C to fahrenheit.
	*/
	fmt.Println("This program to convert Celcius to Fahrenheit.")
	c := cToFah(12)
	fmt.Println(c)
}
