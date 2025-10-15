//Copy the file into backup.txt.

package main
import (
	"fmt"
	"os"
	"io"
)

func main(){
	src,_:=os.Open("output.txt")
	defer src.Close()
	dst,_:=os.Create("copy.txt")
	defer dst.Close()

	_,err:=io.Copy(dst,src)
	if err!=nil{
		fmt.Println("Error: ",err)
		return
	}
	fmt.Println("File Copied!")

}