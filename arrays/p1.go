package main
import "fmt"

func main(){
	var arr=[5]int {2,4,6,8,10}
	sum:=0
	for _,val:=range arr{
		sum=sum+val
	}
	fmt.Println(sum)
}