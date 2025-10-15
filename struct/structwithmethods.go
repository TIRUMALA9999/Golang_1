package main
import "fmt"
type Rectangle struct{
	length,width int
}

func (r Rectangle) Area() int{
	return r.length * r.width
}

func main(){
	rect:=Rectangle{length:10,width:5}
	fmt.Println("Area:", rect.Area())

}
