package main

import (
	"errors"
	"fmt"
)

// Custom error type
type MathError struct {
	Op     string
	A, B   int
}

// Implement error interface
func (e *MathError) Error() string {
	return fmt.Sprintf("cannot %s %d and %d", e.Op, e.A, e.B)
}

// safeDiv safely divides two numbers
func safeDiv(a, b int) (int, error) {
	if b == 0 {
		// Return *MathError instead of generic error
		return 0, &MathError{Op: "divide", A: a, B: b}
	}
	return a / b, nil
}

func main() {
	_, err := safeDiv(5, 0)
	if err != nil {
		fmt.Println("Error:", err)  //Here err satisfies error interface. so it calls err.Error()
		                            //so fmt.Sprintf return here.

		// Try to extract the underlying *MathError
		var me *MathError
		if errors.As(err, &me) {
			fmt.Printf("Operands were: %d and %d\n", me.A, me.B)
		}
	}

}


