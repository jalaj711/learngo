package main

func ___recur(i, j, m, n int, grid [][]int) {
	if i == m-1 && j == n-1 {
		grid[i][j] = -grid[i][j]
		return
	}
	cost := 1000000
	if i < m-1 {
		if grid[i+1][j] >= 0 {
			___recur(i+1, j, m, n, grid)
		}
		cost = -grid[i+1][j]
	}
	if j < n-1 {
		if grid[i][j+1] >= 0 {
			___recur(i, j+1, m, n, grid)
		}
		cost = min(cost, -grid[i][j+1])
	}
	grid[i][j] = -grid[i][j] - cost
}

func minPathSum(grid [][]int) int {
	___recur(0, 0, len(grid), len(grid[0]), grid)
	return -grid[0][0]
}
