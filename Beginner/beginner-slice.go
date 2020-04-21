package main

import (
	"fmt"
)

func change(x []int) {
	x[3] = 2
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}

func main() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Println("Before Change...")
	printSlice(mySlice)
	fmt.Println("##########################################")
	fmt.Printf("Before. Cap: %d, length: %d\n", cap(mySlice), len(mySlice))
	mySlice = append(mySlice, -12)
	fmt.Printf("After. Cap: %d, length: %d\n", cap(mySlice), len(mySlice))
	fmt.Println(mySlice)
	mySlice = append(mySlice, 5)
	fmt.Printf("After after. Cap: %d, length: %d\n", cap(mySlice), len(mySlice))
	fmt.Println(mySlice)
	anotherSlice := make([]int, 4)
	fmt.Printf("A new slice with 4 elements: ")
	printSlice(anotherSlice)
	println(cap(anotherSlice))
}
