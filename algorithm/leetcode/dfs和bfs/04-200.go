/*
	给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
	0103
	0105面试遇到
	0105 再刷(也不知道什么原因 最后结果不对，折腾了40分钟)
*/

// m、n表示网格的长宽
var m, n int

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m = len(grid)
	n = len(grid[0])
	res := 0
	// （1）从左到右从上到下，遍历表格
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, x, y int) {
	// （1）四个方向
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	// （2）标记已经走过的位置
	grid[x][y] = '0'
	// （3）遍历四个方向
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		// 边界条件
		if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] == '1' {
			dfs(grid, a, b)
		}
	}
}


