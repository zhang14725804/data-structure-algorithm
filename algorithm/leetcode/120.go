package leetcode

import (
	"fmt"
	"math"
)

/*
	Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

	For example, given the following triangle

	[
		[2],
		[3,4],
		[6,5,7],
		[4,1,8,3]
	]
	The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

	Note:
	Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

	ps::动态规划和滑动数组的方法
*/
// Leetcode120 动态规划方法
func Leetcode120() int {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	n := len(triangle)
	// 二维数组
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			// 声明一个比较大的值
			dp[i][j] = math.MaxInt64
			// 左右边界
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+triangle[i][j])
			}
			if j < i {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+triangle[i][j])
			}
		}
	}
	// 声明一个比较大的值
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	fmt.Println(res)
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
