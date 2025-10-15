/*Design a ValidationError with:

Field string

Msg string

Use it in a function validateUser(name string, age int) error:

if name is empty → "name cannot be empty"

if age < 0 → "age cannot be negative"

Call it with ("", -5) and print both field + message using errors.As.*/




package main

import (
	"errors"
	"fmt"
)

// Custom error type for validation
type ValidationError struct {
	Field string
	Msg   string
}

// Implement error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// validateUser applies simple validation rules
func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{Field: "name", Msg: "cannot be empty"}
	}
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "cannot be negative"}
	}
	return nil
}

func main() {
	// Call with invalid values
	err := validateUser("", -5)
	if err != nil {
		fmt.Println("Error:", err)

		// Extract structured error
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Msg)
		}
	}
}
