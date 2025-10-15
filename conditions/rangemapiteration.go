package main
import "fmt"
func main(){
	Students:=map[string]int{"Alice": 90,"Bob":80}  //declaring map
	for key,value:=range Students{
		fmt.Println(key,"Scored",value)
	}
}