package main
import (
	"fmt"
	"strconv"
)
func main() {
	isNew:= "tj"
	newString, err := strconv.ParseBool(isNew)
	if err != nil {
		fmt.Println("Failed")
	} else {
		if (newString) {
			fmt.Println("isNew")
		}else {
			fmt.Println("isNotNew")
		}
	}
}
