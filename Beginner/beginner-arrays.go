package main

import "fmt"

func main() {
	myArray := [4]int{1, 2, 3, 4}
	twoDarray := [2][2]int{{1, 3}, {1, 22}}
	threeDarray := [3][3][3]int{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, {{2, 5, 3}, {4, 3, 6}, {3, 6, 5}}, {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
	arrayZero := []int{1, 2, 3, 4}

	fmt.Println(myArray[2])
	fmt.Println(twoDarray[0][1])
	fmt.Println(threeDarray[1][2][2])
	fmt.Println(arrayZero[2])

	for _, number := range threeDarray {
		fmt.Printf("%d", number)
	}
	fmt.Println()
}
