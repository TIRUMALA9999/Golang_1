//You can name return values in the function signature.
//They act like variables → can be returned without explicitly writing them.

package main
import "fmt"

func rectangle(length,width int) (area int, perimeter int){
	area=length*width
	perimeter=2*(length+width)
	return
}

func main(){
	a,p:=rectangle(4,6)
	fmt.Println("Area:", a, "Perimeter:", p)
}