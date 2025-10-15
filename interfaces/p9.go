//Interface Composition (Solution)


package main
import "fmt"

// Define Reader interface
type Reader interface {
	Read() string
}

// Define Writer interface
type Writer interface {
	Write(data string)
}

// File struct implements both
type File struct{}

func (f File) Read() string {
	return "Reading file..."
}

func (f File) Write(data string) {
	fmt.Println("Writing:", data)
}

// Composed interface (inherits from Reader + Writer)
type ReadWriter interface {
	Reader
	Writer
}

// Function that works with ReadWriter
func useReadWriter(rw ReadWriter) {
	fmt.Println(rw.Read())
	rw.Write("Hello.txt")
}

func main() {
	f := File{}
	useReadWriter(f)
}
