
/*
	硬编码，无甚技巧
*/
func isValidSudoku(board [][]byte) bool {
	st := make([]bool, 9)

	for i := 0; i < 9; i++ {
		st = make([]bool, 9)
		for j := 0; j < 9; j++ {
			// 判断行 😅
			if board[i][j] != '.' {
				// 字符的1-9，变成0-8
				t := board[i][j] - '1'
				if st[t] == true {
					return false
				}
				// 标记
				st[t] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		st = make([]bool, 9)
		for j := 0; j < 9; j++ {
			// 判断列 😅
			if board[j][i] != '.' {
				// 字符的1-9，变成0-8
				t := board[j][i] - '1'
				if st[t] == true {
					return false
				}
				//
				st[t] = true
			}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			st = make([]bool, 9)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					// 判断每个九宫格 😅
					if board[i+x][j+y] != '.' {
						// 字符的1-9，变成0-8
						t := board[i+x][j+y] - '1'
						if st[t] == true {
							return false
						}
						st[t] = true
					}
				}
			}
		}
	}
	return true
}