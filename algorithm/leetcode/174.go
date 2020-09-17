/*
	todo：三个思路（回溯，递归，动态规划）
	回溯和递归超时，思路还是可以了解下
	思路：动态规划
*/
func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])
	// 用 dp[i][j] 表示从 (i, j) 到达终点所需要的最小生命值。
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	if dungeon[n-1][m-1] > 0 {
		dp[n-1][m-1] = 1
	} else {
		dp[n-1][m-1] = -dungeon[n-1][m-1] + 1
	}

	for i := 0; i <= n; i++ {
		dp[i][m] = INT_MAX
	}
	for i := 0; i <= m; i++ {
		dp[n][i] = INT_MAX
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}
			// 向下还是向右
			dp[i][j] = MinInt(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if dp[i][j] <= 0 {
				dp[i][j] = 1
			}
		}
	}
	return dp[0][0]
}

// todo：动态规划优化。dp 数组在更新第 i 行的时候，我们只需要第 i+1 行的信息，只需要一维数组即可
func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])

	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = INT_MAX
	}

	if dungeon[n-1][m-1] > 0 {
		dp[n-1] = 1
	} else {
		dp[n-1] = -dungeon[n-1][m-1] + 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}
			// 向下还是向右
			dp[j] = MinInt(dp[j], dp[j+1]) - dungeon[i][j]
			if dp[j] <= 0 {
				dp[j] = 1
			}
		}
	}
	return dp[0]
}