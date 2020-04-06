package main
import (
	"fmt"
	"time"
)

func main(){
	current:= time.Now()
	fmt.Println(current)
	fmt.Println(current.AddDate(0,1,0))
	fmt.Println(current.AddDate(1,0,0))
	fmt.Println(current.AddDate(-1,0,0))
	newTime:= current.Add(10* time.Hour)
	fmt.Println(newTime)

}