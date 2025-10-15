
package main
import "fmt"
func main(){
	students:=make(map[string]int)
	students["Alice"]=90
	students["Bob"]=80
	students["Alice"]=95
	fmt.Println(students)
}

