package main
import "fmt"
func main(){
	t:=[]int{3,4,21,5,6}
	t=append(t,44,55,66)
	fmt.Println("Length: ",len(t))
	fmt.Println("Capacity: ",cap(t))
}