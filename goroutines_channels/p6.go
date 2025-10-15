/*Make a channel of type int. Start a goroutine that sends 10, 20, 30 into the channel. In 
main, receive and print all 3 values.
👉 Expected output:
10
20
30 */

package main
import "fmt"
func main(){
	ch:=make(chan int)
	go func(){
		ch <- 10
		ch <- 20
		ch <- 30
		close(ch)
	}()

	for msg:= range ch {
		fmt.Println(msg)
	}
}