package main

import (
	"fmt"
)

func main() {
	/*
		Logic for Leap year
	    1. If the year get divisible by 4 or 100 or 400 it is considered as leap year.
	    2. Else Not a leap Year.
	*/
	var year int
	fmt.Print("Enter the Year:: ")
	fmt.Scan(&year)

	if year%4 == 0 || year%100 == 0 || year%400 == 0 {
		fmt.Println("Well this is a Leap Year !!")
	} else {
		fmt.Println("This is not a Leap Year..!!")
		}

}
