//Creating or Overwriting a existing file.

package main
import (
	"fmt"
	"os"
)

func main(){
	content:="Hello,How are you man!\n"
	err:=os.WriteFile("output1.txt",[]byte(content),0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File written successfully")
}

/*🧩 1. os.WriteFile

It’s a standard library function from the os package.

Purpose: create (or overwrite) a file and write data into it in one shot.

It’s shorthand for: open → write → close (all done internally).*/

