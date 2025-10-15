package main

import (
    "fmt"
    "time"
    "runtime"
)

func main() {
    fmt.Println("Hello, Golang!")
    fmt.Println("My name is Tirumala.")

    today := time.Now()
    fmt.Println("Today’s date is", today.Format("2006-01-02"))

    fmt.Println("I’m running on", runtime.GOOS)
    fmt.Println("Go version:", runtime.Version())
}
