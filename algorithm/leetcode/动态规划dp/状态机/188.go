/*
 	买卖股票的最佳时机 IV
	给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
	设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
	注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

*/

/*
	方法1：动态规划（状态机）

	dp[i][k][0] // 第i天，至多进行k次交易，目前没有(0/1)持有股票
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
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			// 第i天未持有：前一天未持有（rest）、前一天持有（第i天sell）
			dp[i][k][0] = MaxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			// 第i天持有：前一天持有（rest）、前一天未持有（第i天buy）
			dp[i][k][1] = MaxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][max_k][0]
}

/*
	188变种：你最多可以完成 任意 笔交易
	分两种情况（传入的 k 值可以任意大，导致 dp 数组太大）
		（1）k>n/2。和k为无穷一样处理
		（2）k<=n/2。常规处理
*/