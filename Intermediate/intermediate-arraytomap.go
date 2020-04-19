package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Define an array and map.
	myArray := [4]int{1, 2, 3, 4}
	myMap := make(map[string]int)

	// Function to convert array to map
	lengthArray := len(myArray)
	for i := 0; i < lengthArray; i++ {
		// strconv.Itoa > Convert index to string.
		fmt.Printf("%s ", strconv.Itoa(i))
		myMap[strconv.Itoa(i)] = myArray[i]
	}
	fmt.Println()

	// Print the map
	count := 0
	for key, value := range myMap {
		fmt.Printf("%s: %d\n", key, value)
		count++
	}
}
