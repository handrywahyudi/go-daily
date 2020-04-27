// The example in this subsection will search the provided text for two variations of a given string and replace it with another string

package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	var s [3]string
	s[0] = "1 b 3"
	s[1] = "11 b B 14 1 1"
	s[2] = "a 2 b B -5"

	parse, err := regexp.Compile("[bB]")

	if err != nil {
		fmt.Printf("Error compiling RE: %s\n", err)
		os.Exit(-1)
	}

	for i := 0; i < len(s); i++ {
		temp := parse.ReplaceAllString(s[i], "C")
		fmt.Println(temp)
	}
}
