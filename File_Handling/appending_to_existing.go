package main
import (
	"fmt"
	"os"
)

func main(){
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    fmt.Println("Error:", err)
	return
    }
    defer file.Close()

    if _, err := file.WriteString("Appending a new line\n"); err != nil {
	    fmt.Println("Error:", err)
	}

}