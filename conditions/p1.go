//Write a program to check if a number is positive, negative, or zero.

package main
import "fmt"
func main(){
	x:=0
	if x>0{
		fmt.Println("Number is positive")
	}else if x==0 {
		fmt.Println("Number is zero")
	}else{
		fmt.Println("Number is negative")
	}
}