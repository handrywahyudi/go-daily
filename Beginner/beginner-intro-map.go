package main

import (
	"fmt"
)

func main() {
	// number := make(map[string]int)
	// number["one"] = 1
	// number["two"] = 2
	// number["three"] = 3
	// number["four"] = 4
	// number["five"] = 5
	// number["six"] = 6

	number := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	// delete(number, "one")
	if result, ok := number["seven"]; ok {
		fmt.Println(result, ok)
	} else {
		fmt.Println("The element is not found.")
	}

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
	}

	if el, ok := elements["Be"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// How to find the smallest number
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var n, smallest int
	n = x[0]
	for _, number := range x {
		if number < n {
			n = number
			smallest = number
		}
	}

	fmt.Println(x)
	fmt.Println("The smallest number is ", smallest)

}
