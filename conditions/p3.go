//Use range to print each letter of "Golang" on a new line.

package main
import "fmt"
func main(){
	for i,ch:= range "Golang"{
		fmt.Printf("Index: %d, Char: %c\n",i,ch) //Here \n is used to print each letter in
		                                         // the  new line.
	}
}