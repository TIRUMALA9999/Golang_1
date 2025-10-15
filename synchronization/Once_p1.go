//Use sync.Once to initialize a database connection message only once, even if 5 goroutines try.

package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initDB() {
	fmt.Println("Database connection initialized")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d trying to init DB...\n", id)
			once.Do(initDB) // initDB runs only once
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
	time.Sleep(time.Second)
}


/*🔹 Why sync.Once is Useful

For database connections, logging setup, config file loading, etc.

Prevents accidental double initialization in concurrent code.*/