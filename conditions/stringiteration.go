package main
import "fmt"
func main(){
	t:="Teja"
	for i,ch:=range t{
		fmt.Println(i,"  ",ch)  //if we use like this we get character numbers as output. So we 
		                        // need to format it.
	}

	for i,ch:= range t{
		fmt.Printf("Index:%d,Char:%c\n",i,ch)
	}
}