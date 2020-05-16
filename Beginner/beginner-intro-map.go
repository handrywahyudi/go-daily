package main

import (
	"fmt"
)

func main() {
	number := make(map[string]int)
	number["one"] = 1
	number["two"] = 2
	number["three"] = 3
	number["four"] = 4
	number["five"] = 5
	number["six"] = 6

	delete(number, "one")
	fmt.Println(number)
}
