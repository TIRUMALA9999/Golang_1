package main
import "fmt"

func main(){
	var age int =25
	var name string ="Tirumala"
	isStudent :=true
	const pi=3.14159

	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	fmt.Println("Student?",isStudent)
	fmt.Println("Pi Value:",pi)

	//Type Conversion
	var x int= 10
	var y float64= 5.5
	sum:=float64(x)+y
	fmt.Println("Sum: ",sum)
}