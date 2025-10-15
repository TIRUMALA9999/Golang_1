/*A goroutine is a lightweight thread managed by the Go runtime.
Created using the go keyword before a function call.
Runs concurrently (not necessarily in parallel, but can be if multiple CPU cores are available).*/


package main
import (
    "fmt"
    "time"
)

func printMsg(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg, i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go printMsg("Goroutine") // runs concurrently
    printMsg("Main")  //main goroutine
}
