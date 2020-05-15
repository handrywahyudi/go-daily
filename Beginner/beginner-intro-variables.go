package main

import (
	"fmt"
)

var dogsName string = "dog"

func dog() {
	fmt.Println("i have a", dogsName)
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

	const name string = "handry"
	name = "wahyudi"
	fmt.Println(name)
}
