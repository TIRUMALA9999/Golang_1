/*🟢Send & Receive
Create a channel, send a single string "Hello from Goroutine" into it from a goroutine, and 
print it in main.
👉 Expected output:
Hello from Goroutine */
package main
import "fmt"

func main(){
	ch:=make(chan string)
	go func(){
		ch <- "Hello from Goroutine"
	}()

	fmt.Println(<- ch)
}