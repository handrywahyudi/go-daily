package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide one argument.")
		os.Exit(1)
	}
	filename := arguments[1]

	fileinfo, err := os.Lstat(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if fileinfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(filename, "is a symbolik link.")
		realpath, err := filepath.EvalSymlinks(filename)
		if err == nil {
			fmt.Println("Path:", realpath)
		}
	}
}
