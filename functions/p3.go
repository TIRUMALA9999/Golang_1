//Write a function isEvenOdd(n int) (bool, string) → returns a bool and a message (“Even”/“Odd”).

package main
import "fmt"
func isEvenOdd(n int) (bool,string){
	if n%2==0{
		return true, "Even"
	}
	return false, "Odd"
}

func main(){
	fmt.Println(isEvenOdd(5))
}