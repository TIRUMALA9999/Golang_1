package main
import "fmt"

func main(){
	x:=1
	if x>10{
		fmt.Println("x is larger than 10")
	}else if x==10{
		fmt.Println("x is equal to 10")
	}else{
		fmt.Println("x is less than 10")
	}
}