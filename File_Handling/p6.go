//Read a text file line by line and count how many lines it has.

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){
	file,err:= os.Open("output.txt")
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	count:=0
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		count+=1
	}
	fmt.Println("Number of lines: ",count)
	if err := scanner.Err(); err != nil {
        fmt.Println("Scanner error:", err)
	}
}