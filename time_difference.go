package main
import (
	"fmt"
	"time"
)

func main() {
	firstTime:= time.Date(2019,4,6,13,32,0,0,time.UTC)
	secondTime:= time.Date(2019,5,6,13,32,0,0,time.UTC)
	differenceTime:= secondTime.Sub(firstTime)
	fmt.Printf("Difference %v", differenceTime)
}