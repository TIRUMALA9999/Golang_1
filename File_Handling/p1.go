//Write a program that creates a file log.txt and writes "Log started".

package main
import (
	"fmt"
	"os"
)

func main(){
	content:="Log started"
	err:=os.WriteFile("log.txt",[]byte(content),0644)
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("File created and contents written successfully")
	
	//Reading file contents
	data, _ := os.ReadFile("log.txt")
	fmt.Println(string(data))
}