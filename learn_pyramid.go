
package main
import "fmt"

func main () {
	var rows int
	var k int =0
	fmt.Print("Enter number of Rows: ")
	fmt.Scan(&rows)

	for i :=1; i <= rows; i++ {
		k = 0
		for spaces :=1; spaces <= rows -i; spaces ++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if(k == 2*i -1){
				break
			}
		}
		fmt.Println("")
	}
}