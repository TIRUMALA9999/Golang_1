/*🔹 3. recover() → Catch a Panic

recover is used inside a deferred function.

It lets you catch the panic, so your program doesn’t crash.*/

package main

import "fmt"

func safeDivision(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	return a / b // panic if b == 0
}

func main() {
	fmt.Println("Start")
	fmt.Println(safeDivision(10, 2))
	fmt.Println(safeDivision(5, 0)) // panic handled
	fmt.Println("End")              // program continues
}


/*👉 Without recover(), program would have crashed. With recover(), it continues safely.*/

/*🔹 4. When to Use?

Error handling (error) → for expected problems (e.g., file not found, invalid input).

Panic/Recover → for unexpected, unrecoverable situations (e.g., corrupted memory, programmer mistakes).

Best practice:

Library functions → return error (don’t panic).

Main program / critical system → may use panic for fatal errors.

recover is often used in servers to keep them alive even if one request handler panics.



⚡ Telugu + English Mix

panic ante → program immediate ga crash ayye laaga stop cheyyadam.

recover ante → panic ni catch chesi program ni continue cheyyadam.

Idi usually defer function lo vaadali.*/