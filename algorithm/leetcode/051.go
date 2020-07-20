/*
	n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

	给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

	每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

	思路：全排列问题，要考虑对角线（正反两个方向）
*/
var (
	n int
	// 
	col,dg,udg []bool
	path []string
	ans [][]string
)
func solveNQueens(_n int) [][]string {
	n = _n
	col = make([]bool,n)
	dg,udg = make([]bool,n*2),make([]bool,n*2)
	// 有问题
	path = make([]string,n)
	for i := 0; i < n; i++ {
		path[i] = "."
	}
	dfs(0)
	return ans
}

func dfs(u int){
	if u == n{
		ans = append(ans，path)
		return 
	}

	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i]{
			col[i] , dg[u-i+n] , udg[u+i] = true,true,true
			// 有问题
			path[u][i] = "Q"
			dfs(u+1)
			col[i] , dg[u-i+n] , udg[u+i] = false,false,false
			path[u][i] = "."
		}
	}
}