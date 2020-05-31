package main

import (
    "fmt"
    "flag"
    "os"
    "strings"
)

func main() {
    minusA := flag.Bool("a", false, "a")
    minusS := flag.Bool("s", false, "s")

    flag.Parse()
    flags := flag.Args()
    if len(flags) == 0 {
        fmt.Println("Please provide an argument!.")
        os.Exit(1)
    }

    file := flags[0]
    fountIt := false

    path := os.Getenv("PATH")
    pathslice := strings.Split(path, ":")
    for _, directory := range pathslice {
       fullpath := directory + "/" + file
       fileInfo, err := os.Stat(fullpath)
       if err == nil {
           mode := fileInfo.Mode()
           if mode.IsRegular() {
               if mode&0111 != 0 {
                   fountIt = true
                   if *minusS == true {
                       os.Exit(0)
                   }
                   if *minusA == true {
                       fmt.Println(fullpath)
                   }else{
                       fmt.Println(fullpath)
                       os.Exit(0)
                   }
               }
           }
       }
    }
    if fountIt == false {
         os.Exit(1)
   }
}
