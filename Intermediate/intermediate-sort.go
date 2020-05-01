package main

import (
	"fmt"
	"sort"
)

type myStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]myStructure, 0)
	a := myStructure{"muhammad", 170, 95}
	mySlice = append(mySlice, a)
	a = myStructure{"handry", 160, 85}
	mySlice = append(mySlice, a)
	a = myStructure{"wahyudi", 175, 75}
	mySlice = append(mySlice, a)

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">:", mySlice)

}
