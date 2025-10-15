package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// worker function
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// if cancellation/timeout signal received
			fmt.Printf("Worker %d stopped: %v\n", id, ctx.Err())
			return
		case j, ok := <-jobs:
			if !ok {
				return // no more jobs
			}
			fmt.Printf("Worker %d started job %d\n", id, j)
			time.Sleep(2 * time.Second) // simulate work
			fmt.Printf("Worker %d finished job %d\n", id, j)
			results <- j * 2
		}
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Create a context with timeout (3 sec)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	// Send 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers in background
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
