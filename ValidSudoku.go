package leetcode

func isValidSudoku(board [][]byte) bool {
	var valid map[byte]bool

	for _, line := range board {
		valid = make(map[byte]bool)

		for _, char := range line {
			if char == '.' {
				continue
			}

			if valid[char] {
				return false
			}

			valid[char] = true
		}
	}

	for i := range board {
		valid = make(map[byte]bool)

		for j := range board[i] {
			if board[j][i] == '.' {
				continue
			}

			if valid[board[j][i]] {
				return false
			}

			valid[board[j][i]] = true
		}
	}

	for i := 0; i < 9; i++ {
		valid = make(map[byte]bool)

		for j := 0; j < 9; j++ {
			if board[j/3+(i/3)*3][j%3+(i%3)*3] == '.' {
				continue
			}

			if valid[board[j/3+(i/3)*3][j%3+(i%3)*3]] {
				return false
			}

			valid[board[j/3+(i/3)*3][j%3+(i%3)*3]] = true
		}
	}

	return true
}
