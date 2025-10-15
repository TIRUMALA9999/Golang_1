/*Basic pool (unordered results)

Use the worker pool template to square numbers 1..30 with N=4.

Print jobID and output as they finish.*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// worker squares numbers
func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// simulate some work
		time.Sleep(time.Millisecond * 200)
		output := j * j
		results <- fmt.Sprintf("Worker %d finished job %d → %d", id, j, output)
	}
}

func main() {
	const numJobs = 30
	const numWorkers = 4

	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	var wg sync.WaitGroup

	// Start N=4 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send 30 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Close results once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results (unordered)
	for r := range results {
		fmt.Println(r)
	}
}
