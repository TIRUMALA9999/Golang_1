//Extension to the p8.go

/* Two Goroutines + Select
Create two goroutines:
One sends "Ping" after 1 second.
Another sends "Pong" after 2 seconds.
Use select in main to print whichever message arrives first.
👉 Expected output:
Ping
Pong */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// Goroutine 1 → Ping
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "Ping"
	}()

	// Goroutine 2 → Pong
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Pong"
	}()

	// Receive 2 messages
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}


/* Same single channel

Rendu goroutines okate channel ki send chesthe → receiver side ki just messages ochche order prakaaram print avutayi.

Ante → "Ping" first, "Pong" next (because delays different).

Idi kuda correct ga panichesthundi, but who sent what ani differentiate cheyyadaniki channel id separate ga levu.*/