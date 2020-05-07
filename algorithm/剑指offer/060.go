/*
	todo：又是动态规划问题😅
	（1）状态标识
	（2）状态如何计算
	（3）边界

	todo：难的很，很难
*/
func getMaxValue(grid [][]int) int {
	n,m := len(grid),len(grid[0])
	// 从1开始避免处理数组越界问题0
	f := make([][]int,n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int,m+1)
	}
	// todo：他么代码居然有问题
	for i:=1; i<=n; i++{
		for j:=1; j<=m; j++{
			f[i][j] = max(f[i-1][j], f[i][j-1]+grid[i-1][j-1])
		}
	}
	return f[n][m]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}