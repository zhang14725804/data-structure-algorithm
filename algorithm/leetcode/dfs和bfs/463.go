type pair struct {
	x, y int
}

// 四个方向（上下左右）岛屿问题这里总记不住 😅😅😅
var dir4 = []pair{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

/*
	岛屿问题 😅😅😅 😅😅😅 😅😅😅
	1. 定义四个方向【上下左右】
	2. m，n为方格长宽
	3. 遍历每行每列
	4. 遇到岛屿，进行递归
		a. 遇到边界ans++
		b. 已经走过的跳过，标记已经走过
		c. 遍历上下左右四个方向
*/
func islandPerimeter(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		// 边界情况：可以遍历每个陆地格子，看其四个方向是否为边界或者水域，如果是，将这条边的贡献（即 1）加入答案 ans 中即可。
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			ans++
			return
		}

		// 已经走过
		if grid[x][y] == 2 {
			return
		}
		// 标记已经走过
		grid[x][y] = 2
		// 继续下一个循环
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}
	for i, row := range grid {
		for j, col := range row {
			// 如果是岛屿
			if col == 1 {
				dfs(i, j)
			}
		}
	}
	return
}