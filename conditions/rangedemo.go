package main
import "fmt"
func main(){
	nums:=[]int{10,20,30}
	for index,value:=range nums{
		fmt.Println(index,value)
	}

	//ignoring index

	for _,value:=range nums{
		fmt.Println(value)
	}

	//On string runes
	for i,ch:=range "Go"{
		fmt.Printf("Index:%d,Char:%c\n",i,ch)
	}
}

