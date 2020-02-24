package main 

import "fmt"

func main () {
	// var myInt = 16
	myInt := 16
	// var val, ok = "bishwajit", true
	var val, _ = "bishwajit", true
	fmt.Println("myInt:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val:", val)
	// fmt.Println("ok:", ok)
}