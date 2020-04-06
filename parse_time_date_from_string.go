package main

import (
	"fmt"
	"time"
)

func main() {
	str:= "2020-04-06T14:17:15.349Z"
	layout:= "2006-01-02T15:04:05.000Z"

	t,err := time.Parse(layout, str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t.String())
	}

}
