package main

// import "fmt"

func numDecodings(s string) int {
	size := len(s)
	dp := make([]int, size)
	if s[size-1] != '0' {
		dp[size-1] = 1
	}
	for i := size - 2; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}
		dp[i] = dp[i+1]
		if (s[i] == '1') || (s[i] == '2' && s[i+1] <= '6') {
			if i == size-2 {
				if s[i+1] == '0' {
					dp[i] = 1
				} else {
					dp[i] = 2
				}
			} else {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

// func main() {
// 	fmt.Println(numDecodings("10"))
// 	fmt.Println(numDecodings("226"))
// 	fmt.Println(numDecodings("06"))
// 	fmt.Println(numDecodings("2102"))
// 	fmt.Println(numDecodings("1516"))
// 	fmt.Println(numDecodings("1016"))
// }
