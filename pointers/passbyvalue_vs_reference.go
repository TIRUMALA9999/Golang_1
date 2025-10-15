package main
import "fmt"

func incrementValue(x int){
	x=x+1	
}

func incrementPointer(x *int){
	*x=*x+1	
}

func main(){
	a:=5
	incrementValue(a)
	fmt.Println(a)
	incrementPointer(&a)
	fmt.Println(a)
}

//Shows why pointers are useful for modifying values inside functions.