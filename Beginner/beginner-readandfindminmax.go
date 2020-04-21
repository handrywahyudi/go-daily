package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Write a Go program that keeps reading integers until you give the number 0 as input,
// then it prints the minimum and maximum integer in the input.
func main() {
	reader := bufio.NewReader(os.Stdin)
	empAge := []int{}

	looping := true
	for looping {
		fmt.Print("Give me an age, except '0' : ")
		str, _ := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		i, _ := strconv.Atoi(str)

		if i == 0 {
			looping = false
		} else {
			empAge = append(empAge, i)
			// fmt.Println(myNumberSlice)
		}
	}

	fmt.Println(empAge)
	var n, biggest, smallest int
	for _, number := range empAge {
		if number > n {
			n = number
			biggest = n
		}
	}

	for _, number := range empAge {
		if number < n {
			n = number
			smallest = number
		}
	}

	fmt.Printf("The oldest in this organization is : %v years old.\n", biggest)
	fmt.Printf("The youngest in this organization is : %v years old.\n", smallest)
}
