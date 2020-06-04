package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	overwrite := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()
	flags := flag.Args()

	if len(flags) < 2 {
		fmt.Println("Please provide 2 argument!")
		os.Exit(1)
	}

	source := flags[0]
	destination := flags[1]
	fileInfo, err := os.Stat(source)
	if err == nil {
		mode := fileInfo.Mode()
		if mode.IsRegular() == false {
			fmt.Println("We only support regular file.")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error reading: ", source)
		os.Exit(1)
	}

	newDestination := destination
	destInfo, err := os.Stat(newDestination)
	if err == nil {
		mode := destInfo.Mode()
		if mode.IsDir() {
			justTheName := filepath.Base(source)
			newDestination = destination + "/" + justTheName
		}
	}

	destination = newDestination
	destInfo, err = os.Stat(destination)
	if err == nil {
		if *overwrite == false {
			fmt.Println("The destination already exists.")
			os.Exit(1)
		}
	}

	os.Rename(source, destination)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
