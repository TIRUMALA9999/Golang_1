package main

import "fmt"

func main() {
	// Create channel
	ch := make(chan int)

	// Sender goroutine
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // close after sending
	}()

	// Receiver (main goroutine)
	for val := range ch {
		fmt.Println(val)
	}
}
