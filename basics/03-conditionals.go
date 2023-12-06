package main

import "fmt"

func main() {
	a := 10
	b := 9
	if a > b {
		fmt.Printf("%d > %d", a, b)
	} else if a == b {
		fmt.Printf("%d = %d", a, b)
	} else {
		fmt.Printf("%d < %d", a, b)
	}
	fmt.Println()
}
