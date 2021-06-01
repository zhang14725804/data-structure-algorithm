/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
*/
var (
	n                        int
	colUsed, dgUsed, udgUsed []bool
	board                    [][]byte
	ans                      [][]string
)

func solveNQueens(_n int) [][]string {
	n = _n
	// 列，斜线，反斜线
	colUsed = make([]bool, n)
	dgUsed = make([]bool, n*2)
	udgUsed = make([]bool, n*2)

	board = make([][]byte, n)
	ans = make([][]string, 0)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	backtrack(0)
	return ans
}

// 用row来记录当前遍历到棋盘的第几层了
func backtrack(row int) {
	if row == n {
		str := make([]string, 0)
		for i := 0; i < n; i++ {
			str = append(str, string(board[i]))
		}
		ans = append(ans, str)
		return
	}
	for col := 0; col < n; col++ {
		if !colUsed[col] && !dgUsed[row-col+n] && !udgUsed[row+col] {
			colUsed[col], dgUsed[row-col+n], udgUsed[row+col] = true, true, true
			board[row][col] = 'Q'
			backtrack(row + 1)
			colUsed[col], dgUsed[row-col+n], udgUsed[row+col] = false, false, false
			board[row][col] = '.'
		}
	}
}