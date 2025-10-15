package main
import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, _ := os.Open("example.txt")
	defer src.Close()

	dst, _ := os.Create("copy.txt")
	defer dst.Close()

	_, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println("Copy error:", err)
		return
	}
	fmt.Println("File copied!")
}
