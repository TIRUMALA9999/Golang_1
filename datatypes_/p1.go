//Create variables for: your name (string), age (int), height in meters (float64), 
// student status (bool). Print them.

package main

import (
	"fmt"
)

func main(){
	var name = "Tirumala"    //You need not to mention string here. it will understand automatically
	var age int= 25
	var height float64= 5.5
	isStudent:= true

	fmt.Println("Name: ",name)
	fmt.Println("age: ",age)
	fmt.Println("height: ",height)
	fmt.Println("student? : ",isStudent)
}