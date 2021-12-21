/*
	m*n 个单元格，中间有若干【障碍物】从左上角走到右下角，只能从左到右或者从上到下走，共有几条路径、
*/

/*
	方法1：递归+判断障碍物 😅😅😅
*/
var m, n int
var obstacleGrid [][]int
var visited map[string]int

func uniquePathsWithObstacles(_obstacleGrid [][]int) int {
	obstacleGrid = _obstacleGrid
	m, n = len(obstacleGrid), len(obstacleGrid[0])
	visited = make(map[string]int, 0)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	return dfs(0, 0)
}
func dfs(x, y int) int {
	// 😅 base case，递归出口
	if x == m-1 && y == n-1 {
		return 1
	}

	// 😅 我是在这里判断障碍物的，这里判断有点迟了

	// 向左、向下的走法
	right, bottom := 0, 0
	key := fmt.Sprintf("%v", x+1) + "@" + fmt.Sprintf("%v", y)
	if val, ok := visited[key]; !ok {
		// 😅 判断是否越界和【障碍物】
		if x+1 < m && obstacleGrid[x+1][y] == 0 {
			right = dfs(x+1, y)
		}
	} else {
		right = val
	}

	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y+1)
	if val, ok := visited[key]; !ok {
		// 😅 判断是否越界和【障碍物】
		if y+1 < n && obstacleGrid[x][y+1] == 0 {
			bottom = dfs(x, y+1)
		}
	} else {
		bottom = val
	}

	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y)
	visited[key] = right + bottom
	return visited[key]
}

/*
	动态规划
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// dp[i][j] ：表示从（0 ，0）出发，到(i, j) 有dp[i][j]条不同的路径。
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 😅 dp数组初始化
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	// 从左到右从上到下遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 递推公式
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
	方法2：动态规划  question 😅😅😅😅
	（1）状态表示：包括集合（所有从起点走到[i][j]的路径）、属性（数量）
	（2）状态计算：最后一步往下，最后一步往右
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	// 二维数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 😅 障碍物
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 😅 第一行或者第一列
			if i == 0 && j == 0 {
				dp[i][j] = 1
			}
			// 😅 不是第一列，最后一步往下走
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			// 😅 不是第一行，最后一步往右走
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}
