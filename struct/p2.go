//Create a Circle struct with a method Area() that returns area (πr²).
package main
import "fmt"
type Circle struct{
	radius float64
}

func (c Circle) Area() float64{
	return 3.14*c.radius*c.radius
}

func main(){
	c1:=Circle{radius:5}
	fmt.Println(c1.Area())
}