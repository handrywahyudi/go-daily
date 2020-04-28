package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t1(8)
	x3 := 1

	st1 := reflect.ValueOf(&x1).Elem()
	st2 := reflect.ValueOf(&x2).Elem()
	st3 := reflect.ValueOf(&x3).Elem()

	typeOfX1 := st1.Type()
	typeOfX2 := st2.Type()
	typeOfX3 := st3.Type()

	fmt.Printf("X1 type %s.\n", typeOfX1)
	fmt.Printf("X2 type %s.\n", typeOfX2)
	fmt.Printf("X3 type %s.\n", typeOfX3)

	fmt.Println(st1)
	fmt.Println(st2)
	fmt.Println(st3)

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
}
