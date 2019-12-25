package leetcode

import "fmt"

/*
	m*n 个单元格，中间有若干障碍物从左上角走到右下角，只能从左到右或者从上到下走，共有几条路径
*/
// leetcode63 动态规划
func Leetcode063() int {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	// 二维数组（todos：我这么开辟二维数组是不是有点蠢）
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 障碍物
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 左上角
			if i == 0 && j == 0 {
				dp[i][j] = 1
			}
			// 不是第一列
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			// 不是第一行
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[n-1][m-1]
}
