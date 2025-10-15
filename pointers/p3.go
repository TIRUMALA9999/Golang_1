//Write a program where you pass a slice pointer to a function and append a new element inside it.

package main
import "fmt"

func appeend(p *[]int){
	*p=append(*p,2)
}

func main() {
    p1 := []int{3, 4, 5}
    fmt.Println("Before:", p1)

    appeend(&p1) // pass slice pointer
    fmt.Println("After :", p1)
}