/* 👉 Define an interface Shape with a method Area() float64.
Implement it for:

Square (with side field)

Rectangle (with width, height fields) */

package main
import "fmt"
type Shape interface{
	Area() float64
}

type Square struct{
	side float64
}

func (s Square) Area() float64{
	return s.side*s.side
}

type Rectangle struct{
	length float64
	width float64
}

func (r Rectangle) Area() float64{
	return r.length*r.width
}

func printArea(s Shape) float64{
	return s.Area()
}

func main(){
	s1:=Square{side:4}
	r1:=Rectangle{length:5,width:4}
	fmt.Println("Square Area: ",printArea(s1))
	fmt.Println("Rectangle Area: ",printArea(r1))
}