package main
import "fmt"

type Engine struct{
	horsepower int
}

type Car struct{
	brand string
	Engine     //struct embedding
}

func main(){
	c:=Car{brand:"Toyota",Engine:Engine{horsepower:150}}
	fmt.Println(c.brand,"with HP:",c.Engine.horsepower)
}