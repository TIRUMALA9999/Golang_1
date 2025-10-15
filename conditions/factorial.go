//Find the factorial of a number (e.g., 5 → 120).

package main
import "fmt"
func main(){
	num:=5
	fact:=1
	if num==0{
		fmt.Println("0!= 1")
	}else if num<0{
		fmt.Println("The number is negative")
	}else{
		for i:=2;i<=num;i++{
			fact=fact*i
		}
		fmt.Println(fact)
	}
}