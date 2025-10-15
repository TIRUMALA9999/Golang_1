//Given a slice of numbers, calculate their sum using range.

package main
import "fmt"
func main(){
	num:=[]int{2,35,4,23,76,8,94,46,78,4}
	sum:=0
	for _,val:= range num{
		sum=sum+val
	}
	fmt.Println(sum)
}