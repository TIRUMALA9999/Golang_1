//Write a function minMax(nums ...int) (int, int) → returns smallest & largest from arguments.

package main
import "fmt"

func minMax(nums ...int)(int,int){
	min:=nums[0]
	max:=nums[0]
	for _,num:=range nums{
		if num<min{
			min=num
		}
		if num>max{
			max=num
		}
	}
	return min,max
}


func main(){
	min,max:=minMax(10,20,100,9,1,2,3,44)
	fmt.Println("Minimum: ",min,"Maximum: ",max)
}