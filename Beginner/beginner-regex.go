package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("handry", "Handry Wahyudi")
	fmt.Println(match)

	parse, err := regexp.Compile("[Hh]andry")
	if err != nil {
		fmt.Printf("Error compiling: %s", err)
	} else {
		fmt.Println(parse.MatchString("Handry Wahyudi"))
		fmt.Println(parse.MatchString("handry Wahyudi"))
		fmt.Println(parse.MatchString("H andry Wahyudi"))
		fmt.Println(parse.ReplaceAllString("Handry Wahyudi", "HANDRY"))
	}
}
