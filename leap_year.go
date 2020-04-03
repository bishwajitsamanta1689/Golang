package main

import (
	"fmt"
)

func main() {
	/*
		Logic for Leap year
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
