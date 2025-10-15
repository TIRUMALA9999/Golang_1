/*Backpressure check

Use jobs := make(chan Job, 2) and momentarily pause consumers 
to see producers block once the buffer fills. */

package main

import (
	"fmt"
	"time"
)

type Job struct {
	id int
}

func main() {
	// Small buffer (capacity = 2)
	jobs := make(chan Job, 2)

	// Producer goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Producer trying to send Job %d\n", i)
			jobs <- Job{id: i} // blocks if channel is full
			fmt.Printf("Producer sent Job %d\n", i)
		}
	}()

	// Pause consumer for a while to let producer get stuck
	time.Sleep(2 * time.Second)

	// Consumer starts reading
	for j := range jobs {
		fmt.Printf("Consumer got Job %d\n", j.id)
		time.Sleep(1 * time.Second) // simulate slow consumer
		if j.id == 5 {
			break
		}
	}
}
