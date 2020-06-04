package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// func walkFunc(path string, info os.FileInfo, err error) error {
// 	fileinfo, err := os.Stat(path)
// 	if err != nil {
// 		return err
// 	}

// 	mode := fileinfo.Mode()
// 	if mode.IsDir() || mode.IsRegular() {
// 		fmt.Println(path)
// 	}
// 	return nil
// }

func main() {

	// Define flag arguments
	minS := flag.Bool("s", false, "Sockets")
	minP := flag.Bool("p", false, "Pipes")
	minSL := flag.Bool("sl", false, "Symbolic Link")
	minD := flag.Bool("d", false, "Directores")
	minF := flag.Bool("f", false, "Files")

	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("No argument that given.")
		os.Exit(1)
	}

	printAll := false

	if *minS && *minP && *minSL && *minD && *minF {
		printAll = true
	}

	if *minS || *minP || *minSL || *minD || *minF {
		printAll = true
	}

	Path := flags[0]

	// Define function walkFunc
	walkFunc := func(path string, info os.FileInfo, err error) error {
		fileinfo, err := os.Stat(path)
		if err != nil {
			return err
		}

		if printAll {
			return nil
		}

		mode := fileinfo.Mode()
		if mode.IsRegular() && *minF {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *minD {
			fmt.Println(path)
			return nil
		}

		fileinfo, _ = os.Lstat(path)
		if fileinfo.Mode()&os.ModeSymlink != 0 {
			if *minSL {
				fmt.Println(path)
				return nil
			}
		}

		if fileinfo.Mode()&os.ModeNamedPipe != 0 {
			if *minP {
				fmt.Println(path)
				return nil
			}
		}

		if fileinfo.Mode()&os.ModeSocket != 0 {
			if *minS {
				fmt.Println(path)
				return nil
			}
		}

		return nil
	}

	err := filepath.Walk(Path, walkFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
