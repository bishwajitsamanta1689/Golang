package main 

import "fmt"

func main () {
	ages:= map[string]int{}
	ages["Bishwajit"] = 50

	if ages["Bishwajit"] < 30 {
		fmt.Println("You cant vote!!")
	} else if ages["Bishwajit"] >51 {
		fmt.Println("You need to Plan for retirement!")
	} else {
		fmt.Println("Exception Caused")
	}
}