/*Unbuffered channel → sender blocks until receiver is ready.

Buffered channel → has capacity; sender doesn’t block until it’s full.*/

package main
import "fmt"
func main() {
	ch := make(chan int, 2) // buffer size 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}
