//A function is a block of code that can be called by name.

package main
import "fmt"

func greet(name string) string{
	return "Hello " + name
}

func main(){
	fmt.Println(greet("Tirumala"))
}