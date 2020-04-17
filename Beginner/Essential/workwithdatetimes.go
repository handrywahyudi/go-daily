package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(1989, time.November, 11, 10, 10, 0, 0, time.UTC)
	fmt.Printf("My Birthday is : %s\n", t)

	now := time.Now()
	fmt.Printf("Now is : %s\n", now)

	fmt.Println("The Month is :", t.Month())
	fmt.Println("The Day is : ", t.Day())
	fmt.Println("The Weekday is :", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	tomorrowFromNow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow from now is %v, %v, %v, %v\n", tomorrowFromNow.Weekday(), tomorrowFromNow.Month(), tomorrowFromNow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 1989"
	fmt.Println("Tomorrow is : ", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is : ", tomorrow.Format(shortFormat))
}
