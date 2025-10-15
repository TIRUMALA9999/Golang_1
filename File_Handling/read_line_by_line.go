package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	file,err:=os.Open("example.txt")
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
	defer file.Close()

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println("Line: ",scanner.Text())
	}

	if err:=scanner.Err();err!=nil{
		fmt.Println("Error: ",err)
	}
}