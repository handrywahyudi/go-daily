package main

import (
	"fmt"
)

func main() {
	i, j := 72, 3421
	p := &i         // pointer to i
	fmt.Println(*p) // read i through the pointer
	*p = 27         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // pointer to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

}
