/*
 	买卖股票的最佳时机 IV
	k=any integer
*/

/*
	方法1：动态规划（状态机）
	穷举所有状态(天数i，交易次数k，当前持有状态，三种情况穷举)
*/
func maxProfit(max_k int, prices []int) int {
	// 初始化
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, max_k+1)
		for j := 0; j < max_k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	// 对天数进行穷举
	for i := 0; i < n; i++ {
		// 对交易次数进行穷举
		for k := max_k; k >= 1; k-- {
			// 处理base case
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = MaxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = MaxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][max_k][0]
}