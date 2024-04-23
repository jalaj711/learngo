package main

func recur(i, j, m, n int, dp [][]int) {
	if i == m-1 && j == n-1 {
		return
	}
	dp[i][j] = 0
	if i < m-1 {
		if dp[i+1][j] < 0 {
			recur(i+1, j, m, n, dp)
		}
		dp[i][j] += dp[i+1][j]
	}
	if j < n-1 {
		if dp[i][j+1] < 0 {
			recur(i, j+1, m, n, dp)
		}
		dp[i][j] += dp[i][j+1]
	}
	return
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = -1
		}
	}
	dp[m-1][n-1] = 1
	recur(0, 0, m, n, dp)
	return dp[0][0]
}

func main() {
	uniquePaths(2, 2)
}
