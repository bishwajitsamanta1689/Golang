package main 

import "fmt"

func main () {
	birthdays := map[string]string {
		"bishwajit": "16/02/1989",
		"Uday" : "1/1/1000",
		"Samanta": "1/1/1",
	}
	delete (birthdays, "Samanta")
	fmt.Println("Birthdays:", birthdays)

	ages := map[string]int {}
	ages["Bishwajit"] = 30
	ages["Uday"] = 65
	ages["Samanta"] = 67


	fmt.Println("age of Bishwajit:", ages["Samanta"])
}