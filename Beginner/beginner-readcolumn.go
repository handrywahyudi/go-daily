package main

import (
	"fmt"
	"strings"
)

func main() {
	var s [3]string
	s[0] = "1 2 3"
	s[1] = "11 22 33 44 55 66"
	s[2] = "-2 3 -5 3 -5"

	column := 1
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		fmt.Println(data)
		if len(data) >= column {
			fmt.Println((data[column-1]))
		}
	}
}
