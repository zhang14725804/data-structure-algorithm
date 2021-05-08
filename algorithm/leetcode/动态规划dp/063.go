/*
	m*n 个单元格，中间有若干障碍物从左上角走到右下角，只能从左到右或者从上到下走，共有几条路径、

	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
*/

/*
	方法1：递归+判断障碍物 😅😅😅
*/
var obstacleGrid [][]int
var m, n int
var visited map[string]int

func uniquePathsWithObstacles(_obstacleGrid [][]int) int {
	obstacleGrid = _obstacleGrid
	visited = make(map[string]int)
	m, n = len(obstacleGrid)-1, len(obstacleGrid[0])-1
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	return dfs(0, 0)
}

func dfs(x, y int) int {
	if x == m && n == y {
		return 1
	}
	var n1, n2 int
	key := fmt.Sprintf("%v", x+1) + "@" + fmt.Sprintf("%v", y)
	if val, ok := visited[key]; !ok {
		if x+1 <= m && obstacleGrid[x+1][y] == 0 {
			n1 = dfs(x+1, y)
		}
	} else {
		n1 = val
	}

	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y+1)
	if val, ok := visited[key]; !ok {
		if y+1 <= n && obstacleGrid[x][y+1] == 0 {
			n2 = dfs(x, y+1)
		}
	} else {
		n2 = val
	}
	key = fmt.Sprintf("%v", x) + "@" + fmt.Sprintf("%v", y)
	visited[key] = n1 + n2
	return visited[key]

}

/*
	方法2：动态规划  question 😅😅😅😅
	（1）状态表示：包括集合（所有从起点走到[i][j]的路径）、属性（数量）
	（2）状态计算：最后一步往下，最后一步往右
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	n, m := len(obstacleGrid), len(obstacleGrid[0])
	// 二维数组（todos：我这么开辟二维数组是不是有点蠢）
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
			// 😅 左上角
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
