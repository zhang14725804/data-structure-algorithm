import (
	"fmt"
	"math"
)

/*
	给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
	相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
	如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	思路：动态规划（我居然因为这个坐标问题困惑）
		（1）状态表示：包括集合（所有从起点走到[i][j]的路径）、属性（Max，所有路径上的数的和的最大值）
		（2）状态计算：最后一步从左上角下来，最后一步从右上角下来

*/
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	// 二维数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	// 第一个点
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			// 声明一个比较大的值
			dp[i][j] = math.MaxInt64
			// 左上角下来[i-1][j-1]
			if j > 0 {
				dp[i][j] = compare(dp[i][j], dp[i-1][j-1]+triangle[i][j], true)
			}
			//右上角[i-1][j]
			if j < i {
				dp[i][j] = compare(dp[i][j], dp[i-1][j]+triangle[i][j], true)
			}
		}
	}
	// 声明一个比较大的值
	res := math.MaxInt64
	// 遍历最后一行，取最大值
	for i := 0; i < n; i++ {
		res = compare(res, dp[n-1][i], true)
	}
	fmt.Println(res)
	return res
}

// todo：这个骚操作，不懂
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	// 二维数组，只需要两行就搞定
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
	}
	// 第一个点
	dp[0][0] = triangle[0][0]
	// 所有的点都除以2，也就是 "& 1",骚操作
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			// 声明一个比较大的值
			dp[i&1][j] = math.MaxInt64
			// 左上角下来[i-1][j-1]
			if j > 0 {
				dp[i&1][j] = compare(dp[i&1][j], dp[i-1&1][j-1]+triangle[i][j], true)
			}
			//右上角[i-1][j]
			if j < i {
				// 这里数组越界
				dp[i&1][j] = compare(dp[i&1][j], dp[i-1&1][j]+triangle[i][j], true)
			}
		}
	}
	// 声明一个比较大的值
	res := math.MaxInt64
	// 遍历最后一行，取最大值
	for i := 0; i < n; i++ {
		res = compare(res, dp[n-1&1][i], true)
	}
	return res
}