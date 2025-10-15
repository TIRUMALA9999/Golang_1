package main 

import (
	"runtime"
	"fmt"
)

func main(){
	fmt.Println("My name is Tirumala Teja")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Version())
}