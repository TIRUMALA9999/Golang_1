//Write a function swap(a *int, b *int) that swaps two numbers using pointers.

package main
import "fmt"

func swap(a *int,b *int){
	*a,*b=*b,*a
}

func main(){
	x:=7
	y:=3
	fmt.Println("Before swap:", x, y)
    swap(&x, &y)
    fmt.Println("After swap:", x, y)
}
