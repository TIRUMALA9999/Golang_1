//Use switch to print the name of a month given its number (1 = Jan … 12 = Dec).

package main
import "fmt"
func main(){
	num:=5

	switch num{
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("February")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	default:
		fmt.Println("month name")

	}
}