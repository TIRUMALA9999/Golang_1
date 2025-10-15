/*A channel is used to communicate safely between goroutines.

Think of it as a pipe: one goroutine sends data, another receives it.

Created with make(chan Type).*/

package main
import "fmt"

func worker(ch chan string) {
    ch <- "Work done!" // send to channel
}

func main() {
    ch := make(chan string)
    go worker(ch)

    result := <-ch // receive from channel
    fmt.Println(result)
}
