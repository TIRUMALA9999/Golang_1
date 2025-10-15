//Go functions can return more than one value. Common in error handling.

package main
import "fmt"
func divide(a, b int) (int, bool) { //multiple return values int, bool(true/false)
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
    result, ok := divide(10, 2)
    fmt.Println("Result:", result, "Success?", ok)
}
