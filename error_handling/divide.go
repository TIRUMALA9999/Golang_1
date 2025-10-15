package main

import (
	"errors"
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	// 1) Basic: ok path
	ans, err := div(10, 2)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Println("ok:", ans)
}