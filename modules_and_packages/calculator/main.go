/*Create a module calculator with two packages:

add → function Add(a, b int) int

mul → function Multiply(a, b int) int
In main.go, import both and print results.*/


package main

import (
	"fmt"
	"calculator/add"
	"calculator/mul"
)

func main() {
	sum := add.Add(5, 3)
	product := mul.Multiply(5, 3)

	fmt.Println("5 + 3 =", sum)
	fmt.Println("5 * 3 =", product)
}