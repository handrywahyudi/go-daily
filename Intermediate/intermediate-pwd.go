package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	argument := os.Args

	pwd, err := os.Getwd()
	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println("Error : ", err)
	}

	if len(argument) == 1 {
		return
	}

	if argument[1] != "-P" {
		return
	}

	fileinfo, err := os.Lstat(pwd)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		realpath, err := filepath.EvalSymlinks(pwd)
		if err == nil {
			fmt.Println(realpath)
		}
	}
}
