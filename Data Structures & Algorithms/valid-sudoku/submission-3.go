func isValidSudoku(board [][]byte) bool {
	for i := range board {
		if !isValidRow(board[i]) {
			return false
		}

		if !isValidColumn(board, i) {
			return false
		}

		if !isValidQuad(board, i) {
			return false
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	m := make(map[byte]struct{})

	for i := range row {
		if row[i] == '.' {
			continue
		}

		if _, ok := m[row[i]]; ok {
			return false
		}

		m[row[i]] = struct{}{}
	}

	return true
}

func isValidColumn(board [][]byte, colNum int) bool {
	m := make(map[byte]struct{})

	for i := range board {
		if board[i][colNum] == '.' {
			continue
		}

		if _, ok := m[board[i][colNum]]; ok {
			return false
		}

		m[board[i][colNum]] = struct{}{}
	}

	return true
}

func isValidQuad(board [][]byte, quadNum int) bool {
	m := make(map[byte]struct{})

	addI := quadNum / 3

	addJ := quadNum % 3

	for i := 0 + addI * 3; i < addI * 3 + 3; i++ {
		for j := 0 + addJ * 3; j < addJ * 3 + 3; j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, ok := m[board[i][j]]; ok {
				return false
			}

			m[board[i][j]] = struct{}{}
		}
	}

	return true
}
