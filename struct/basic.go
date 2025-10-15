package main
import "fmt"
type Student struct{
	name string
	age int
	grade string
}

func main(){
	s1:=Student{name:"Alice",age:20,grade:"A"}
	fmt.Println(s1.name,s1.age,s1.grade)
}