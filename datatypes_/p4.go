package main

import "fmt"

func main(){
	var a int =10
	var b int =3
	sum:=a+b
	difference:=a-b 
	product:=a*b 
	division:=float64(a)/float64(b)
	fmt.Println("Sum: ",sum)
	fmt.Println("difference: ",difference)
	fmt.Println("product: ",product)
	fmt.Println("floor dividion: ",division)
}