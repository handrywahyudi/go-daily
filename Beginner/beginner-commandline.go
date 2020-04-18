package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Get the command line arguments
	// arguments := os.Args
	// for i := 0; i < len(arguments); i++ {
	// 	fmt.Println(arguments[i])
	// }

	// Get the arguments and sum that all
	arguments := os.Args
	sum := 0
	for i := 0; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum = sum + temp
	}
	fmt.Println("Sum is : ", sum)
}
