package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	overwrite := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()
	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("Please provide 2 argument!")
		os.Exit(1)
	}

	fmt.Println(flags[0])
	fmt.Println(flags[1])

	fmt.Println(overwrite)

}
