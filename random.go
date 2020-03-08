package main
import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1)
	var f int 
	fmt.Print("Enter a number in Number Range 1-100:: ")
	_, err := fmt.Scanf("%d", &f)
  if err != nil {
		fmt.Println(err)
	}
	
	// value := rand.Intn(100)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println("Printing a Random Integer from 0-100: ", rand.Intn(f))
}