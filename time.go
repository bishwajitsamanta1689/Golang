package main
import (
	"fmt"
	"time"
)
func main(){
	current:= time.Now()
	fmt.Println(current)
	fmt.Println(current.Format("02-01-2006 15:04:05"))
}