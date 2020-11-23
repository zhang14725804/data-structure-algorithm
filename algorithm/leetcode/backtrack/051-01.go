/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	方法：基于集合的回溯
*/
var ans [][]string
var queues []int // 互斥锁
var cols *Set    // 列是否有皇后
var diag1 *Set   // 斜线是否有皇后
var diag2 *Set   // 反斜线是否有皇后
var n int

func solveNQueens(_n int) [][]string {
	n = _n
	cols = NewSet()
	diag1 = NewSet()
	diag2 = NewSet()
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
	} else {
		for i := 0; i < n; i++ {
			// 每一列是否存在
			if cols.Contains(i) {
				continue
			}
			// 左上到右下方向的斜线：行下边与列下标只差相等
			d1 := row - i
			if diag1.Contains(d1) {
				continue
			}
			// 右上到左下方向的斜线：行下表与列下表之和相等
			d2 := row + i
			if diag2.Contains(d2) {
				continue
			}

			queues[row] = i
			cols.Insert(i)
			diag1.Insert(d1)
			diag2.Insert(d2)
			backtrack(row + 1)

			// 状态初始化
			queues[row] = -1
			cols.Remove(i)
			diag1.Remove(d1)
			diag2.Remove(d2)
		}
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