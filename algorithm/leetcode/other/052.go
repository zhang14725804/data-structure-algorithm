var ans, n int

// 列，主副对角线
var col, dg, udg []bool

func totalNQueens(_n int) int {
	n = _n
	col = make([]bool, n)
	dg, udg = make([]bool, n*2), make([]bool, n*2)
	backtrack(0)
	return ans
}

func backtrack(u int) {
	if u == n {
		ans++
		return
	}
	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i] {
			col[i], dg[u-i+n], udg[u+i] = true, true, true
			backtrack(u + 1)
			col[i], dg[u-i+n], udg[u+i] = false, false, false
		}
	}
}