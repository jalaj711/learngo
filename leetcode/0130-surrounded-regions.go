package main

func __dfs(board *[][]byte, startx, starty, m, n int, vis *[][]bool) {
	(*vis)[startx][starty] = true
	if startx > 0 && (*vis)[startx-1][starty] == false && (*board)[startx-1][starty] == 'O' {
		__dfs(board, startx-1, starty, m, n, vis)
	}
	if startx < m-1 && (*vis)[startx+1][starty] == false && (*board)[startx+1][starty] == 'O' {
		__dfs(board, startx+1, starty, m, n, vis)
	}
	if starty < n-1 && (*vis)[startx][starty+1] == false && (*board)[startx][starty+1] == 'O' {
		__dfs(board, startx, starty+1, m, n, vis)
	}
	if starty > 0 && (*vis)[startx][starty-1] == false && (*board)[startx][starty-1] == 'O' {
		__dfs(board, startx, starty-1, m, n, vis)
	}
}
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !vis[i][0] {
			__dfs(&board, i, 0, m, n, &vis)
		}
		if board[i][n-1] == 'O' && !vis[i][n-1] {
			__dfs(&board, i, n-1, m, n, &vis)
		}
	}
	for i := 1; i < n-1; i++ {
		if board[0][i] == 'O' && !vis[0][i] {
			__dfs(&board, 0, i, m, n, &vis)
		}
		if board[m-1][i] == 'O' && !vis[m-1][i] {
			__dfs(&board, m-1, i, m, n, &vis)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && vis[i][j] != true {
				board[i][j] = 'X'
			}
		}
	}
}

//func main() {
//	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
//	solve(board)
//}
