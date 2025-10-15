//Nice 🚀 Let’s move into the Empty Interface (interface{}) — this is the “accept anything”
//  feature of Go. It’s used a lot in logging, JSON parsing, and libraries like fmt.
/*Write a function:

func printAnything(x interface{}) {
    // should print both type and value
}

Requirements:

Accepts any value (int, string, bool, struct, etc.).

Prints its type (%T) and value (%v).

Test it with:

an integer 42

a string "Teja"

a struct Guitar{} from before*/

package main
import "fmt"

type Guitar struct{}

func printAnything(x interface{}){
	fmt.Printf("Type: %T,Value: %v\n",x,x)
}

func main(){
	printAnything(42)
	printAnything("Teja")
	printAnything(Guitar{})
}