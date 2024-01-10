package main

func isValidSudoku(board [][]byte) bool {
	check := map[byte]bool{}
	flag := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if _, pre := check[board[i][j]]; pre {
				flag = true
				break
			}
			check[board[i][j]] = true
		}
		check = map[byte]bool{}
	}
	if flag {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if _, pre := check[board[j][i]]; pre {
				flag = true
				break
			}
			check[board[j][i]] = true
		}
		check = map[byte]bool{}
	}
	if flag {
		return false
	}

	for xoffset := 0; xoffset < 3; xoffset++ {
		for yoffset := 0; yoffset < 3; yoffset++ {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[xoffset*3+i][yoffset*3+j] == '.' {
						continue
					}
					if _, pre := check[board[xoffset*3+i][yoffset*3+j]]; pre {
						flag = true
						break
					}
					check[board[xoffset*3+i][yoffset*3+j]] = true
				}
			}

			check = map[byte]bool{}
		}
	}
	return !flag
}

// func main() {
// 	board := [][]byte{
// 		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
// 	}
// 	isValidSudoku(board)
// }
