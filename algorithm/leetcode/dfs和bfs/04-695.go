/*
	给定一个包含了一些 0 和 1 的非空二维数组 grid 。
	一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
	你可以假设 grid 的四个边缘都被 0（代表水）包围着。
	找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

*/
// 网格长宽
var m, n int

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m = len(grid)
	n = len(grid[0])
	// 遍历网格
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果是岛屿
			if grid[i][j] == 1 {
				res = compare(res, dfs(grid, i, j), true)
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int) int {
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}
	count := 1
	// 标记走过的位置
	grid[x][y] = 0
	// 遍历四个方向
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		// 边界条件
		if a >= 0 && a < m && b >= 0 && b < n && grid[a][b] == 1 {
			count += dfs(grid, a, b)
		}
	}
	return count
}
