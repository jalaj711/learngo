package main

// import "fmt"

func minCost(colors string, neededTime []int) int {
	time := 0
	currentMax := -1
	for i := 0; i < len(colors); i++ {
		if i < len(colors)-1 && colors[i] == colors[i+1] {
			for i < len(colors)-1 && colors[i] == colors[i+1] {
				time += neededTime[i]
				if neededTime[i] > currentMax {
					currentMax = neededTime[i]
				}
				i++
			}
			time += neededTime[i]
			if neededTime[i] > currentMax {
				currentMax = neededTime[i]
			}
			time -= currentMax
			currentMax = -1
		}
	}
	return time
}

// func main() {
// 	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
// 	fmt.Println(minCost("abc", []int{1, 2, 3}))
// 	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
// 	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}))
// }
