// Write a program that removes an element from a slice at index 2.

package main
import "fmt"
func main(){
	t:=[]int{1,2,3,4,5,6,7}
	index:=2
	t=append(t[:index],t[index+1:]...)
	fmt.Println(t)
}

//sum(1,2,3) → you’re giving items one by one.
//sum(values...) → you’re giving the whole slice, but ... unpacks it into individual items.