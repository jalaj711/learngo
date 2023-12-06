package main

import "fmt"

func main() {
	var a int = 10
	var b uint = 0o23424
	var c complex128 = 3587 + 309478i
	d := 90
	const e string = "some string"
	fmt.Println("Value of a is:", a)
	fmt.Println("Value of b is:", b)
	fmt.Println("Value of c is:", c)
	fmt.Println("Value of d is:", d)
	fmt.Println("Value of e is:", e)
}
