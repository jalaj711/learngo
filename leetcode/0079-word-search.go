package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	var _search func(x, y, k int) bool

	_search = func(x, y, k int) bool {
		if k == len(word) {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[k] {
			return false
		}

		temp := board[x][y]
		board[x][y] = '0'

		if _search(x+1, y, k+1) || _search(x-1, y, k+1) ||
			_search(x, y+1, k+1) || _search(x, y-1, k+1) {
			return true
		}

		board[x][y] = temp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if _search(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}

//
//func main() {
//	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
//	//exist(board, "ABCCED")
//	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
//	//exist(board, "SEE")
//	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
//	//exist(board, "ABCB")
//	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
//	exist(board, "ABCESEEEFS")
//	//board := [][]byte{{'A', 'A'}}
//	//exist(board, "AAA")
//	//board := [][]byte{{'A', 'B'}}
//	//exist(board, "AB")
//}
