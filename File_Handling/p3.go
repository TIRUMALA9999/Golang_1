//Read the file line by line and print each line.

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file,err:=os.Open("output.txt")
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err:=scanner.Err();err!=nil{
		fmt.Println("Error: ",err)
	}


}