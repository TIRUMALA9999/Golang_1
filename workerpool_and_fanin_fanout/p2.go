/*Ordered results
Modify the pool so results print in input order 
(hint: send ID, write into a results := make([]int, total), then print after collection).  */

package main

import (
	"fmt"
	"sync"
	"time"
)

// Result struct holds job ID and computed value
type Result struct {
	id    int
	value int
}

// worker squares numbers and sends Result{id, value}
func worker(id int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(time.Millisecond * 200) // simulate work
		results <- Result{id: j, value: j * j}
	}
}

func main() {
	const numJobs = 30
	const numWorkers = 4

	jobs := make(chan int, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start 4 workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs (IDs are just 1..30)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Close results once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Create ordered results slice
	ordered := make([]int, numJobs)

	// Collect into ordered slice
	for r := range results {
		ordered[r.id-1] = r.value
	}

	// Print in input order
	for i, val := range ordered {
		fmt.Printf("Job %d → %d\n", i+1, val)
	}
}
