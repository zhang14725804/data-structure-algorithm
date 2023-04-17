/*
	方法1:动态规划
	😅 😅 😅 没有思路
*/

func numSquares(n int) int {
	// dp 表示最少需要多少个数的平方来表示整数 i
	// dp[0]为边界情况
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := math.MaxInt32
		//
		for j := 1; j*j <= i; j++ {
			min = Min(min, dp[i-j*j])
		}
		//
		dp[i] = min + 1
	}
	return dp[n]
}
