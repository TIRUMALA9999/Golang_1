//Define a Car struct that embeds Engine. Add a method Details() that prints 
// car brand and horsepower.

package main
import "fmt"

type Engine struct{
	horsepower int
}

type Car struct{
	brand string
	Engine
}

func (c Car) Details(){
	fmt.Println(c.brand,c.horsepower)
}

func main(){
	c1:=Car{brand:"Toyota",Engine:Engine{horsepower:155}}
	c1.Details()
}