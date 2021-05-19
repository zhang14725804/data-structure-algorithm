/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	思路：全排列问题，要考虑对角线（正反两个方向）
*/
var (
	n            int
	col, dg, udg []bool
	path         [][]byte
	ans          [][]string
)

func solveNQueens(_n int) [][]string {
	n = _n
	col = make([]bool, n)
	// 斜线，反斜线
	dg, udg = make([]bool, n*2), make([]bool, n*2)
	path = make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}

	backtrack(0)
	return ans
}

func backtrack(u int) {
	if u == n {
		spath := make([]string, 0)
		for i := 0; i < n; i++ {
			// []byte转string
			spath = append(spath, string(path[i][:]))
		}
		ans = append(ans, spath)
		return
	}
	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i] {
			col[i], dg[u-i+n], udg[u+i] = true, true, true
			path[u][i] = 'Q'
			backtrack(u + 1)
			col[i], dg[u-i+n], udg[u+i] = false, false, false
			path[u][i] = '.'
		}
	}
}