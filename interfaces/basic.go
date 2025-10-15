/* Create an interface Animal with a method Sound() string.

Implement it for Dog and Cat.

Print their sounds using the interface. */

package main
import "fmt"

type Animal interface{
	Sound() string
}

type Dog struct{}

func (d Dog) Sound() string{
	return "woof!"
}

type Cat struct{}
func (c Cat) Sound() string{
	return "meow!"
}

func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func main() {
	d := Dog{}
	c := Cat{}

	makeSound(d) // Woof!
	makeSound(c) // Meow!
}