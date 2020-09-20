/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	思路：全排列问题，要考虑对角线（正反两个方向）
*/
var (
	n int
	//
	col, dg, udg []bool
	path         []string
	ans          [][]string
)

func solveNQueens(_n int) [][]string {
	n = _n
	col = make([]bool, n)
	dg, udg = make([]bool, n*2), make([]bool, n*2)
	// 有问题
	path = make([]string, n)
	for i := 0; i < n; i++ {
		path[i] = "."
	}
	dfs(0)
	return ans
}

func dfs(u int) {
	if u == n {
		ans = append(ans, path)
		return
	}

	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i] {
			col[i], dg[u-i+n], udg[u+i] = true, true, true
			// 有问题
			path[u][i] = "Q"
			dfs(u + 1)
			col[i], dg[u-i+n], udg[u+i] = false, false, false
			path[u][i] = "."
		}
	}
}

/*
	方法1：基于集合的回溯

	todo：用set存放列，斜线，反斜线是否有皇后，可以用bool数组代替
*/
var ans [][]string
var queues []int

// 用三个集合cols，diag1，diag2表示列，斜线，反斜线是否有皇后
var cols *Set
var diag1 *Set
var diag2 *Set
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
			// 判断每一列是否存在
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

/*
	方法2：基于位运算的回溯（todo：没看懂😅）
*/