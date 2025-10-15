//Parse a JSON file of students and calculate the average marks.

package main
import (
	"fmt"
	"os"
	"encoding/json"
)

type Student struct{
	Name string
	Marks int
}

func main(){
	file,err:=os.ReadFile("students.json")
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}

	var students []Student
	if err:=json.Unmarshal(file,&students);err!=nil{
		fmt.Println("Error Parsing Json: ",err)
		return 
	}

	if len(students)==0{
		fmt.Println("No students found")
		return
	}

	total:=0
	for _,s:=range students{
		total+=s.Marks
	}
	average:=float64(total)/float64(len(students))

	fmt.Printf("Total Students:%d\n",len(students))
	fmt.Printf("Average Marks: %.2f\n",average)
}