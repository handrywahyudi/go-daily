package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	minTest := flag.Bool("test", false, "Test run!")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 || len(flags) == 1 {
		fmt.Println("Not enough argument")
		os.Exit(1)
	}

	Path := flags[0]
	fmt.Println(Path)
}
