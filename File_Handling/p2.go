//Append "New event logged" to the same file.

package main
import (
	"fmt"
	"os"
)

func main(){
	file,err:=os.OpenFile("log.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("Error: ",err)
	}
	defer file.Close()

	if _,err:=file.WriteString("New Event Logged\n");err!=nil{
		fmt.Println("Error: ",err)
	}

}