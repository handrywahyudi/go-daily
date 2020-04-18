package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	i1, i2, i3 := 12, 15, 3
	intSum := i1 + i2 + i3
	fmt.Printf("Result : %v\n", intSum)

	f1, f2, f3 := 34.2, 32.1, 54.7
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum is : ", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(34.2)
	b2.SetFloat64(32.1)
	b3.SetFloat64(54.7)

	bigSum.Add(&b1, &b1).Add(&bigSum, &b3)
	fmt.Printf("Big Sum is: %.10g\n", &bigSum)

	circleRadius := 15.5
	circumference := circleRadius * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
