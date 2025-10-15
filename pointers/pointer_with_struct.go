package main
import "fmt"
type Student struct {
    name string
    age  int
}

func main(){
	s:=Student{name:"Teja",age:25}
	p:=&s 
	fmt.Println(p.age)
	p.age=21
	fmt.Println(s.age)
}