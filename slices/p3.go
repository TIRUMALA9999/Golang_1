//Given a slice [3, 1, 4, 1, 5, 9], find the maximum element.

package main
import "fmt"
func main(){
	t:=[]int{3,1,4,1,5,9}
	max:=t[0]
	for _,num:=range t{
		if num>max{
			max=num
		}
	}
	fmt.Println(max)
}