package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	min := 0
	max := 0
	total := 0

	fmt.Println(len(os.Args))
	if len(os.Args) > 3 {
		min, _ = strconv.Atoi(os.Args[1])
		max, _ = strconv.Atoi(os.Args[2])
		total, _ = strconv.Atoi(os.Args[3])
	} else {
		fmt.Println("Usage: ", os.Args[0], "MIN MAX TOTAL")
		os.Exit(-1)
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < total; i++ {
		myRand := random(min, max)
		fmt.Print(myRand)
		fmt.Print(" ")
	}

	fmt.Println()
}
