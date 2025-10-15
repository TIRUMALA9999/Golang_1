package main

import (
	"errors"
	"fmt"
	"strconv"
)

// parseNumber converts string → int
func parseNumber(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		// Wrap error with context + %w
		return 0, fmt.Errorf("parseNumber failed: %w", err)
	}
	return n, nil
}

func main() {
	// Intentionally pass bad input
	i, err := parseNumber("hello")
	if err != nil {
		fmt.Println(err)

		// Check if it’s a syntax error (bad number format)
		if errors.Is(err, strconv.ErrSyntax) {
			fmt.Println("reason: invalid number syntax")
		}
	}else{
		fmt.Println(i)
	}
}
