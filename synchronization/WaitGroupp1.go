//Use a WaitGroup to run 5 workers in parallel, each printing "Hello from worker X".

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // one more goroutine
		go func(id int) {
			defer wg.Done() // mark done when finished
			fmt.Printf("Hello from worker %d\n", id)
		}(i)
	}

	wg.Wait() // wait until all goroutines finish
	fmt.Println("All workers finished")
}
