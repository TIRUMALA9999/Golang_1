/*🔹 2. panic(err) → Immediate Stop

panic is like saying "something went terribly wrong, I cannot continue."

It stops the normal flow of the program.

It unwinds the stack → all deferred functions are executed, then program crashes.*/

package main

import "fmt"

func riskyDivision(a, b int) int {
	if b == 0 {
		panic("division by zero") // stop everything
	}
	return a / b
}

func main() {
	fmt.Println("Start")
	fmt.Println(riskyDivision(10, 2))
	fmt.Println(riskyDivision(5, 0)) // panic here
	fmt.Println("End")                // never reached
}
