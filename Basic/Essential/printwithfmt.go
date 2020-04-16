package main

import "fmt"

func main() {
	str1 := "Muhammad"
	str2 := "Handry"
	str3 := "Wahyudi"
	aNumber := 31
	isTrue := true
	//	fmt.Println(str1, str2, str3)

	stringLength, _ := fmt.Println(str1, str2, str3)

	//	if err == nil {
	fmt.Println("String Lenght : ", stringLength)
	//}

	fmt.Printf("Value of a number : %v\n", aNumber)
	fmt.Printf("Value of a isTrue : %v\n", isTrue)
	fmt.Printf("Value of a number as float : %.2f\n", float64(aNumber))

	fmt.Printf("Data Types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)

	myString := fmt.Sprintf("Data Types as var: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)

	fmt.Print(myString)
}
