package main
import "fmt"

func main(){
	x:=10
	p:=&x        // p points to x
	fmt.Println(*p)   // 10 (dereference)
	*p=20              // modify value via pointer
	fmt.Println(x)    //20
}