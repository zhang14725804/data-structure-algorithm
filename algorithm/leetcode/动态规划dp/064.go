/*
	给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
*/

/*
	方法1：动态规划 😄😄😄
	状态定义：设 dp 为大小 m×n 矩阵，其中 dp[i][j] 的值代表直到走到 (i,j) 的最小路径和。
*/
func minPathSum(grid [][]int) int {
	// question 无需新建dp二维数组，直接在原数组操作。机智😄😄😄😄😄😄
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				// 😅左上角
				continue
			} else if i == 0 {
				// 😅左边界
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				// 😅上边界
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				// 😅中间
				grid[i][j] = MinInt(grid[i][j-1], grid[i-1][j]) + grid[i][j]
			}
		}
	}
	return grid[n-1][m-1]
}

/*
	方法2：递归 😄😄😄
	dp[x][y]。这样的话， dp[x][y] = Min（dp[x][y+1] + dp[x+1][y]）+ grid[x][y]。当前点的右边和下边取一个和较小的，然后加上当前点的权值
*/
var grid [][]int
var visited map[string]int // 缓存重叠项
var m, n int

func minPathSum(_grid [][]int) int {
	grid = _grid
	m, n = len(grid)-1, len(grid[0])-1
	visited = make(map[string]int)
	return dfs(0, 0)
}

func dfs(x, y int) int {
	if x == m && y == n {
		return grid[m][n]
	}
	n1, n2 := (1 << 32), (1 << 32)
	// 走右边
	key := fmt.Sprintf("%v", x+1) + "@" + fmt.Sprintf("%v", y)
	if val, ok := visited[key]; !ok {
		// 判断边界
		if x+1 <= m {
			n1 = dfs(x+1, y)
		}
	} else {
		n1 = val
	}
	// 走下边
	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y+1)
	if val, ok := visited[key]; !ok {
		// 判断边界
		if y+1 <= n {
			n2 = dfs(x, y+1)
		}
	} else {
		n2 = val
	}
	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y)
	// 当前点的右边和下边取一个和较小的，然后加上当前点的权值
	val := MinInt(n1, n2) + grid[x][y]
	visited[key] = val
	return val
}