//select lets you wait on multiple channels.

/*select is Go’s way of waiting on multiple channels at once.
If multiple channels are ready, Go picks one randomly.
Adding a default case makes it non-blocking.*/

package main
import "fmt"
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "from ch1" }()
	go func() { ch2 <- "from ch2" }()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
