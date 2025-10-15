/*Problem 7: Type Assertion

👉 Steps:
Store the value "Hello, Go!" inside a variable of type interface{}.
Use a type assertion to check if it’s a string.
If yes, print: It’s a string: Hello, Go!.
Otherwise, print: Not a string. */

package main
import "fmt"
func main(){
	var x interface{}
	x="Hello,Go!"
	val,ok:=x.(string)
	if ok{
		fmt.Println("It is a string",val)
	}else{
		fmt.Println("It is not a string")
	}
}

