package main
import (
	"fmt"
	"strings"

)
func main() {
	greeting := "\t Hello Bishwajit "
	fmt.Printf("%d,%s\n", len(greeting), greeting )
	trimmedValue := strings.TrimSpace(greeting)
	fmt.Printf("%d,%s\n", len(trimmedValue), trimmedValue)
}
