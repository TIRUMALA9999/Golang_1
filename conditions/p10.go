//Given a map of subject → marks, print them in the format:
//Math: 90
//Science: 85
//English: 92

package main
import "fmt"
func main(){
	Student:=map[string]int{"Math":90,"Science":85,"English":92}
	for key,val:= range Student{
		fmt.Println(key,":",val)
	}
}