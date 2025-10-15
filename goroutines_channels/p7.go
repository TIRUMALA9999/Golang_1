/*Buffered Channel
Create a buffered channel with capacity 2. Send 1 and 2 into it without blocking, 
then receive them in main.
👉 Expected output:
1
2*/

package main
import "fmt"
func main(){
	ch:=make(chan int,2)
	go func(){
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for msg:= range ch{
		fmt.Println(msg)
	}
}