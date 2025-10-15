//Write a program that keeps looping until it finds the first multiple of 7 greater than 100.

package main
import "fmt"
func main(){
	i:=1
	for{
		if i%7==0 && i>100{
			fmt.Println("First multiple of 7 greater than 100 is:", i)
			break
		}
		i++
	}
}