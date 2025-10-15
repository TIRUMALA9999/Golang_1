package main
import "fmt"

func main() {
	ch := make(chan string)

	// Sender goroutine
	go func() {
		ch <- "Hello from goroutine" // send
	}()

	// Receiver (main goroutine)
	msg := <-ch // receive
	fmt.Println(msg)
}
