//Use sync.Map to count word frequencies concurrently from multiple goroutines.

/*Normally, a Go map is not safe for concurrent writes. But sync.Map allows multiple goroutines 
to safely update a shared map without needing explicit Mutex locking.
🔹 Task
We’ll:
Split a text into words.
Launch multiple goroutines, each counting words.
Store/update counts in a sync.Map.
Finally, print the word frequencies.*/

package main
import (
	"fmt"
	"sync"
	"strings"
)

func main(){
	var wg sync.WaitGroup
	var m sync.Map

	text:= "go is fun and go is fast but go is also simple and fun"
	words:= strings.Fields(text)

	//worker function
	countWord:= func(word string){
		defer wg.Done()
		val,_:=m.LoadOrStore(word,0)
		m.Store(word,val.(int)+1)
	}

	for _,word:=range words{
		wg.Add(1)
		go countWord(word)
	}

	wg.Wait()

	// print results
	fmt.Println("Word frequencies:")
	m.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key.(string), value.(int))
		return true
	})
}


/*🔹 Explanation

sync.Map → safe concurrent map.

LoadOrStore(key, value) → returns existing value if present, else stores new one.

Store(key, value) → updates the value.

Range → iterate over all key-value pairs.

For each word, we launch a goroutine countWord.

It increments the count in sync.Map.

No need to manually mu.Lock() because sync.Map is already safe.

After all goroutines finish (wg.Wait()), we Range through the map and print frequencies.*/

/*⚡ In short:

sync.Map = concurrent map (safe for multi-goroutine writes).

Very useful in word counting, caching, or storing results from many workers.*/