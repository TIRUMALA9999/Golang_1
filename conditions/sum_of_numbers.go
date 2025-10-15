//Print the sum of numbers from 1 to 100 using a for loop.

package main
import "fmt"
func main(){
	sum:=0
	for i:=1;i<=100;i++{
		sum=sum+i
	}
	fmt.Println(sum)
}