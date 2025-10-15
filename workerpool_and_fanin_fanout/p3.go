/* Cancellation
Start a pool on 1..1000. Cancel with context.WithCancel once 
you observe any result > 300*300, then stop everything cleanly. */

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker squares numbers and respects ctx cancellation
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Stop worker when context is cancelled
			fmt.Printf("Worker %d stopped: %v\n", id, ctx.Err())
			return
		case j, ok := <-jobs:
			if !ok {
				return // no more jobs
			}
			time.Sleep(time.Millisecond * 5) // simulate some work
			results <- j * j
		}
	}
}

func main() {
	const numJobs = 1000
	const numWorkers = 4

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Create cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Collector goroutine: watches results
	go func() {
		for r := range results {
			fmt.Println("Result:", r)
			if r > 300*300 {
				fmt.Println("Condition met! Cancelling all workers...")
				cancel() // 🔴 cancel when square > 90000
				return
			}
		}
	}()

	// Wait for workers, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Block main until results channel fully drains
	for range results {
		// just consume until closed
	}
	fmt.Println("All done (cancelled cleanly).")
}
