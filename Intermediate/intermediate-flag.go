package main

import (
	"flag"
	"fmt"
)

func main() {
	minusA := flag.Bool("a", false, "a")
	minusB := flag.Bool("b", false, "b")
	minusC := flag.Int("c", 0, "an integer")

	flag.Parse()

	fmt.Println("-o", *minusA)
	fmt.Println("-c", *minusB)
	fmt.Println("-k", *minusC)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
