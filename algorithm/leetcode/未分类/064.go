/*
	给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

	动态规划
	状态定义：设 dpdp 为大小 m \times nm×n 矩阵，其中 dp[i][j]dp[i][j] 的值代表直到走到 (i,j)(i,j) 的最小路径和。

*/
func minPathSum(grid [][]int) int {
	// 无需新建dp二维数组，直接在原数组操作，机智
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				// 左上角
				continue
			} else if i == 0 {
				// 左边界
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				// 右边界
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				// 中间
				grid[i][j] = MinInt(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[n-1][m-1]
}