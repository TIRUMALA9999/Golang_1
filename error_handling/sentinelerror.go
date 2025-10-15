/*Define:

var ErrUnauthorized = errors.New("unauthorized")
Write a function checkUser(role string) error.
If role != "admin" → return wrapped ErrUnauthorized.
In main(), call with "guest".
Use errors.Is to check and print "access denied".*/



package main

import (
	"errors"
	"fmt"
)

// Step 1: define sentinel error
var ErrUnauthorized = errors.New("unauthorized")

// Step 2: function that may return wrapped error
func checkUser(role string) error {
	if role != "admin" {
		// wrap sentinel error with context
		return fmt.Errorf("checkUser(%q): %w", role, ErrUnauthorized)
	}
	return nil
}

func main() {
	// Step 3: call with "guest"
	err := checkUser("guest")
	if err != nil {
		fmt.Println("error:", err) // shows wrapped message

		// Step 4: inspect with errors.Is
		if errors.Is(err, ErrUnauthorized) {
			fmt.Println("access denied")
		}
	}
}
