//FizzBuzz Challenge to check the output for numbers 1 → 20.

//For multiples of 3 → print "Fizz", for multiples of 5 → "Buzz", for both → "FizzBuzz".

package main
import "fmt"
func main(){
	for i:=1;i<=20;i++{
		if i%3==0 && i%5==0{
		fmt.Println("FizzBuzz")
	}else if i%3==0{
		fmt.Println("Fizz")
	}else if i%5==0{
		fmt.Println("Buzz")
	}else{
		fmt.Println("The given number is: ",i)
	}
	}	
}