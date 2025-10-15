//Write a goroutine that prints numbers from 1 to 5 while the main goroutine 
// prints letters A–E. Observe interleaving.

package main

import (
	"fmt"
	"time"
)

func numbers(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func letters(s string) {
	for _, val := range s {
		fmt.Println(string(val))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go numbers(5)
	letters("ABCDE")

	//time.Sleep(2*time.Second) // wait for numbers goroutine
}
