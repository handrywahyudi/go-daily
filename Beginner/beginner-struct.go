package main

import (
	"fmt"
	"reflect"
)

// How to define data structure or struct
type information struct {
	Name     string
	Age      int
	Position string
}

func main() {
	// How to add the field of struct.
	employee1 := information{"Mike", 32, "Network Administrator."}
	employee2 := information{"john", 40, "head of infrastructure."}
	employee3 := information{}
	employee3.Name = "Cheik"
	employee3.Age = 34
	employee3.Position = "SRE."

	s1 := reflect.ValueOf(&employee1).Elem()
	s2 := reflect.ValueOf(&employee2).Elem()
	s3 := reflect.ValueOf(&employee3).Elem()
	fmt.Println(s1, s2, s3)
	fmt.Println(s1.Type())
	typeOfT := s1.Type()

	// How to print of struct.
	fmt.Println(employee1, employee2, employee3)

	for i := 0; i < s1.NumField(); i++ {
		f := s1.Field(i)

		fmt.Printf("%d: %s ", i, typeOfT.Field(i).Name)
		fmt.Printf("%s = %v\n", f.Type(), f.Interface())
	}

}
