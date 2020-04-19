package main

import (
	"fmt"
)

func main() {
	// Create a map
	myMap := make(map[string]int)

	// Adding value on a map
	myMap["Mon"] = 1
	myMap["Tue"] = 2
	myMap["Wed"] = 3
	myMap["Thu"] = 4
	myMap["Fri"] = 5
	myMap["Sat"] = 6
	myMap["Sun"] = 7

	// Print the map with key
	fmt.Printf("Wednesday is the %dth day of the week.\n", myMap["Wed"])

	// Check wether the key exists on a map or not.
	_, ok := myMap["Friday"]
	if ok {
		fmt.Printf("The key exists!\n")
	} else {
		fmt.Printf("The key does not exist!\n")
	}

	// How to iterate over all the keys of an existing map.
	count := 0
	for key, value := range myMap {
		count++
		fmt.Printf("%s : %d \n", key, value)
	}

	// How to count elements of map.
	count = 0
	for _, _ = range myMap {
		count++
	}
	fmt.Printf("The map has %d elements.\n", count)

	// How to create the new map, print the value, count the element.
	newMap := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
	}

	count = 0
	for key, value := range newMap {
		count++
		fmt.Printf("%s: %d\n", key, value)
	}
	fmt.Printf("The map has %d elements.\n", count)

}
