package main

import (
    "fmt"
    "flag"
    "os"
    "string"
)

func main() {
    minusA := flang.bool("a", false, "a")
    minusB := flang.bool("b", false, "b")

    flag.parse()
    flags := flag.Args()
    if len(flags) == 0 {
        fmt.Printlns("Please provide an argument!.")
        os.Exit(1)
    }

    file := flags[0]
    foundIt := false

    path := os.Getenv("PATH")
    pathslice := strings.Split(path, ":")
    for _, directory := range pathslice {
       fullpath := directory + "/" + file

    }
}
