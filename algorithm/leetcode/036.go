/*
	判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

	数独部分空格内已填入了数字，空白格用 '.' 表示。

	判断每行每列有没有重复的数字
*/
func isValidSudoku(board [][]byte) bool {
	st:=make([]bool,9)

	// 判断行
	for i := 0; i < 9; i++ {
		st = make([]bool,9)
		for j := 0; j < 9; j++ {
			if board[i][j]!='.'{
				// 字符的1-9，变成0-8
				t:=board[i][j]-'1'
				if st[t] == true{
					return false
				}
				st[t] = true
			}
		}
	}
	// 判断列
	for i := 0; i < 9; i++ {
		st = make([]bool,9)
		for j := 0; j < 9; j++ {
			if board[j][i]!='.'{
				// 字符的1-9，变成0-8
				t:=board[j][i]-'1'
				if st[t] == true{
					return false
				}
				st[t] = true
			}
		}
	}
	// 判断每个九宫格
	for i := 0; i < 9; i+=3 {
		for j := 0; j < 9; j+=3 {
			st = make([]bool,9)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[i+x][j+y]!='.'{
						// 字符的1-9，变成0-8
						t:=board[i+x][j+y]-'1'
						if st[t] == true{
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