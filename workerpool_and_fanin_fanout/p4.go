/*Fan-in

Create 3 channels producing ints at different 
speeds; merge with merge and print the first 20 values. */

package main

import (
	"fmt"
	"time"
)

// generator produces ints with given delay
func generator(name string, delay time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			ch <- i
			fmt.Printf("%s produced %d\n", name, i)
			time.Sleep(delay)
		}
	}()
	return ch
}

// merge combines multiple channels into one
func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	for _, c := range channels {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
		}(c)
	}
	return out
}

func main() {
	// 3 producers at different speeds
	c1 := generator("Fast", 100*time.Millisecond)
	c2 := generator("Medium", 250*time.Millisecond)
	c3 := generator("Slow", 400*time.Millisecond)

	// Merge them
	merged := merge(c1, c2, c3)

	// Print first 20 values
	for i := 0; i < 20; i++ {
		val := <-merged
		fmt.Println("Merged got:", val)
	}
}
