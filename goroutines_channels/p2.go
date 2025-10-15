//Create a channel and send 3 integers into it from one goroutine, and receive them in main.

package main
import "fmt"



func main(){
	ch:=make(chan int)
	go func(){
	ch <- 123
	}()

	msg:= <- ch
	fmt.Println(msg)
}