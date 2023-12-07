package main

import (
	"errors"
	"fmt"
)

// individual declaration of type
func add(a int, b int) int {
	return a + b
}

// only one declaration of type
func mult(a, b int) int {
	return a * b
}

// multi returns
func get_point(a, b int) (int, int) {
	return a * 2, b * 2
}

// named return values
func named_returns(a, b int) (x, y int) {
	x = 8 * a
	y = b * 78
	// naked return
	return
	// named return
	// return x, y
}

// early return, multi return and  error
func divide(a, b int) (ans int, err error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Printf("Sum of 3 and 4: %d", add(3, 4))

	//ignoring a return value
	x, _ := get_point(3, 4)
	fmt.Printf("valkue of x: %d", x)
}
