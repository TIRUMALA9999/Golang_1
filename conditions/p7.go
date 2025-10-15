//Print all even numbers between 1–50 using a for loop.

package main
import "fmt"

func main(){
	for i:=1;i<=50;i++{
		if i%2==0{
			fmt.Println(i)
		}
	}
}