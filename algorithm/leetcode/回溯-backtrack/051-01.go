/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
*/

/*
	方法：回溯

*/
var ans [][]string
var queues []int                    // 互斥锁
var cols = make(map[int]bool, 0)    // 列 是否有皇后
var diag45 = make(map[int]bool, 0)  // 斜线 是否有皇后
var diag135 = make(map[int]bool, 0) // 反斜线 是否有皇后
var n int

func solveNQueens(_n int) [][]string {
	n = _n
	ans = make([][]string, 0)
	// 初始化slice值-1
	queues = make([]int, n)
	for i := 0; i < n; i++ {
		queues[i] = -1
	}
	backtrack(0)
	return ans
}

func backtrack(row int) {
	if row == n {
		board := generate(queues, n)
		ans = append(ans, board)
		return
	}
	for col := 0; col < n; col++ {
		// 排除不合法的选择
		// 每一列是否存在
		if cols[col] {
			continue
		}
		// 左上到右下方向的斜线： 行与列 之差相等
		d1 := row - col
		if diag45[d1] {
			continue
		}
		// 右上到左下方向的斜线： 行与列 之和相等
		d2 := row + col
		if diag135[d2] {
			continue
		}

		// 做选择
		queues[row] = col
		cols[col] = true
		diag45[d1] = true
		diag135[d2] = true

		// 进入下一层决策树
		backtrack(row + 1)

		// 撤销选择
		queues[row] = -1
		cols[col] = false
		diag45[d1] = false
		diag135[d2] = false
	}
}

func generate(queues []int, n int) []string {
	board := make([]string, 0)
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			if j == queues[i] {
				row += "Q"
			} else {
				row += "."
			}
		}
		board = append(board, row)
	}
	return board
}