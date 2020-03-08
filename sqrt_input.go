package main
import (
	"fmt"
	//"strconv"
	"math"
)

func main() {
	// Declare Float Variable
	var f float64

	fmt.Print("Enter a Float Value:: ")
	_, err:= fmt.Scanf("%f", &f)

	if err!= nil {
		fmt.Println(err)
	}
	//fmt.Println("The Number you have inserted ::", strconv.FormatFloat(f, 'f',-1,64))
	fmt.Printf("Square Root Value is: %f \n", math.Sqrt(f))
}