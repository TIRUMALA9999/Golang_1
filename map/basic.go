package main
import "fmt"
func main(){
	students:=map[string]int{
		"Alice":90,
		"Bob":85,
	}
	fmt.Println(students["Alice"])
}