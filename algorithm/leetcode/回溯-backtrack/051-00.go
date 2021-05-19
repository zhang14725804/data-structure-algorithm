var ans [][]string

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	// 初始化每个位置
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	backtrack(board, 0)
	return ans
}

func backtrack(board [][]byte, row int) {
	// 结束条件
	if row == len(board) {
		c := make([]string, 0)
		for i := 0; i < len(board); i++ {
			c = append(c, string(board[i]))
		}
		ans = append(ans, c)
		return
	}

	m := len(board[row])
	for col := 0; col < m; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 选择
		board[row][col] = 'Q'
		// 进入下一层决策树
		backtrack(board, row+1)
		// 撤销选择
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row, col int) bool {
	n := len(board)
	// 列
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 右上方
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 左上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}