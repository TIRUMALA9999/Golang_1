/* 🔹 1. What is Synchronization?
Synchronization ante goroutines madhya coordination.

When multiple goroutines try to read/write same resource (like variable, file, DB) → race conditions vastayi.

Synchronization tools help to avoid data corruption and control execution order. */

/*📝 Why Synchronization?

When multiple goroutines share data, we need tools to coordinate access. Otherwise → race conditions (unpredictable results).

Go provides:

sync.WaitGroup → wait for multiple goroutines to finish.

sync.Mutex → ensure only one goroutine accesses critical section at a time.

sync.Once → run an init function once, even if multiple goroutines call it.

sync.Map → concurrency-safe map.*/

/* 🔹 2. Why do we need it?

Imagine:

2 goroutines are depositing money into the same bank account.

If both update balance at the same time → wrong value vastundi.

We need a way to lock the account → only 1 goroutine updates at a time. */

/* 🔹 3. Synchronization Tools in Go
✅ (a) WaitGroup

Already seen in worker pool.

Helps to wait until goroutines finish.

var wg sync.WaitGroup
wg.Add(1)
go func() {
    fmt.Println("Hello")
    wg.Done()
}()
wg.Wait()

✅ (b) Mutex (Mutual Exclusion Lock)

Only one goroutine can hold the lock at a time.

Prevents race conditions.

package main
import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	count := 0

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()              // lock resource
			count++                // safe update
			mu.Unlock()            // unlock resource
		}()
	}

	wg.Wait()
	fmt.Println("Final count:", count)
}


🔹 Without mu.Lock(), result may be wrong (race condition).
🔹 With mutex → always correct.

✅ (c) RWMutex

A special lock:

Multiple readers can read at same time.

But only one writer at a time.

Good for caches or configs → reads are common, writes are rare.

✅ (d) Channels as Synchronization

Channels themselves can synchronize goroutines.

Example: one goroutine waits until another sends a signal.

done := make(chan bool)

go func() {
    fmt.Println("Work done")
    done <- true
}()

<-done // wait for signal

🔹 4. Real-world Analogy

WaitGroup → Teacher waits until all students submit homework.

Mutex → Only 1 person can use ATM machine at a time.

RWMutex → Many can read a book in library (photocopies), but only 1 person can edit the master copy.

Channel → Like a phone call signal “I’m done, you can continue.”

⚡ In short:
👉 Synchronization = making goroutines work together without conflicts.
👉 Tools: WaitGroup, Mutex, RWMutex, Channels. */


/*🔹 1. sync.WaitGroup

👉 Purpose: Wait until a group of goroutines finish.

Example:
package main
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // tell wg: one more goroutine
		go func(id int) {
			defer wg.Done() // mark goroutine done
			fmt.Println("Worker", id, "done")
		}(i)
	}

	wg.Wait() // block until all Done() called
	fmt.Println("All workers finished")
}


✅ Output:

Worker 1 done
Worker 2 done
Worker 3 done
All workers finished


🔹 Analogy: Teacher waits until all students submit homework.

🔹 2. sync.Mutex

👉 Purpose: Protect a shared resource so only one goroutine can access it at a time.

Example:
package main
import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	count := 0

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()     // lock critical section
			count++       // safe update
			mu.Unlock()   // unlock
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", count)
}


✅ Without mu.Lock(), race condition → wrong count.
✅ With mu.Lock(), always correct.

🔹 Analogy: Only one person can use an ATM machine at a time.

🔹 3. sync.Once

👉 Purpose: Ensure something runs only once (even with multiple goroutines).

Example:
package main
import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	init := func() {
		fmt.Println("Initialized only once!")
	}

	for i := 0; i < 5; i++ {
		go func() {
			once.Do(init) // runs only once
		}()
	}

	// wait a little to see output
	fmt.Scanln()
}


✅ Output:

Initialized only once!


(no matter how many goroutines call it)

🔹 Analogy: Light switch in a room → only the first person turns it ON, others don’t need to.

🔹 4. sync.Map

👉 Purpose: Concurrent map safe for multiple goroutines (thread-safe).
👉 Normal map is not safe for concurrent use.

Example:
package main
import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Store values
	m.Store("name", "Teja")
	m.Store("age", 24)

	// Load value
	if v, ok := m.Load("name"); ok {
		fmt.Println("Name:", v)
	}

	// Range over map
	m.Range(func(key, value any) bool {
		fmt.Println(key, "=", value)
		return true
	})
}


✅ Output:

Name: Teja
name = Teja
age = 24


🔹 Analogy: Like a thread-safe dictionary everyone can write/read at the same time without breaking it.

⚡ Summary
Tool	Use-case	Analogy
WaitGroup	Wait for goroutines to finish	Teacher waits for homework
Mutex	Exclusive lock for shared resource	Only 1 person at ATM
Once	Ensure a function runs only once	First person turns on light switch
Map	Thread-safe map for concurrent read/writes	Shared whiteboard with safety */

/* 5. Other Sync Tools

sync.RWMutex → multiple readers allowed, one writer.

sync.Once → ensure code runs once (singleton pattern).

sync.Cond → condition variables for complex signaling.

sync/atomic → lightweight atomic operations (atomic.AddInt32).

🚀 Summary

Mutex → protect shared data.

WaitGroup → wait for goroutines to finish.

Channels → natural synchronization tool.*/