package main
import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Hello Payel, how are you.. I am fine Payel..  I love You Payel.."
	replaced_value := strings.Replace(greeting,"Payel","Bishwajit",-1)
	fmt.Println(replaced_value)
}
