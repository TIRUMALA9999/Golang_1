//Use a Mutex to safely increment a counter from 100 goroutines and print the final value.

package main
import (
	"sync"
	"fmt"
)

func main(){
	var mu sync.Mutex
	var counter int = 0
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait() // wait until all goroutines finished
	fmt.Println("Final Counter:", counter)
}
