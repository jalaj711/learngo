package main

import "fmt"

func main() {
	// simple for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d;", i)
	}
	fmt.Println()

	// the condition is optional. this creates an infinite loop
	for j := 0; ; j++ {
		if j == 10 {
			break
		}
		fmt.Printf("j=%d;", j)
	}
	fmt.Println()

	// there is no while loop in go but it can be implemented using
	// the `for` keyword itself
	k := 0
	for k < 10 {
		fmt.Printf("k=%d;", k)
		k++
	}
	fmt.Println()

	// the continue keyword works like any other programming language
	for l := 0; l < 10; l++ {
		if l%3 == 0 {
			continue
		}
		fmt.Printf("l=%d;", l)
	}
	fmt.Println()
}
