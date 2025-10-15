// Create a struct Person{name, age} and write a function that updates the age using a pointer.

package main
import "fmt"
type Person struct{
	name string
	age int
}

func (p *Person) update(){
	p.age=21
}

func main() {
    p1 := Person{name: "Teja", age: 25}
    fmt.Println("Before:", p1)

    p1.update() // no need to take &p1 explicitly, Go handles it
    fmt.Println("After :", p1)
}