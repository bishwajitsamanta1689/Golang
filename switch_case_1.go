package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	weekday_weekend := time.Now().Weekday()
	switch weekday_weekend {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend Time")
	default:
		fmt.Println("It's a Weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("This is before Noon")
	default:
		fmt.Println("This is after Noon!")
	}
	whatamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a Bool!!")
		case int:
			fmt.Println("I am a Int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatamI(true)
	whatamI(1)
	whatamI("hey")
}
