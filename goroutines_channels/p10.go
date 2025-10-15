//Use a channel to send squares of numbers (1–5) from one goroutine to another.

package main
import "fmt"
func main(){
	ch:=make(chan int)
	go func(){
		for i:=1;i<=5;i++{
			ch <- i*i
		}
		//close(ch)
	}()

	for j:=1;j<=5;j++{
		fmt.Println(<-ch)
	}
}