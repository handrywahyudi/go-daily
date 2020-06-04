package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	fileinfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileinfo.Mode()
	if mode.IsDir() || mode.IsRegular() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("No argument that given.")
		os.Exit(1)
	}

	Path := flags[0]

	err := filepath.Walk(Path, walkFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
