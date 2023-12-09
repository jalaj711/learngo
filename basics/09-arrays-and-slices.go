package main

import "fmt"

func printArray(arr []int) {

}

func variadic_sum(ints ...int) int {
	n := 0
	// ints becomes a slice
	for i := 0; i < len(ints); i++ {
		n += ints[i]
	}
	return n
}

func main() {
	// basic arrays are fixed size
	var squares [5]int
	for i := 0; i < 5; i++ {
		squares[i] = i * i
	}

	//arrays can also be declared with the := syntax
	strs := [2]string{"hello", "hi"}
	strs[1] = "world"

	//slices
	sqs := squares[1:3]
	// other formats
	// sqs := squares[:3]
	// sqs := squares[3:]
	// sqs := squares[:]

	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))

	// using the make function
	// make(type, len, cap)
	sqs = make([]int, 5, 10)
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))
	// capacity can be skipped
	sqs = make([]int, 5)
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))
	// literal slice
	sqs = []int{1, 2, 3, 4, 5}
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))

	// spread operator
	sum := variadic_sum(sqs...)
	fmt.Printf("sum of sqs=%d\n", sum)

	// append function
	sqs = append(sqs, 6)
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))
	// append is variadic function
	sqs = append(sqs, 7, 8)
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))
	// can also send slices
	sqs = append(sqs, []int{9, 10}...)
	fmt.Printf("Length of sqs=%d;capacity of sqs=%d\n", len(sqs), cap(sqs))

	// 2D slices
	_ = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9},
	}

	// range keyword
	for index, element := range sqs {
		fmt.Printf("sqs[%d]=%d;", index, element)
	}
	fmt.Println()

}
