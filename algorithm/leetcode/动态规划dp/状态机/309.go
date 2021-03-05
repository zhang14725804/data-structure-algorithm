/*
	最佳买卖股票时机含冷冻期

	给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
	设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
		你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
		卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

	k = +infinity with cooldown
	dp[i][k][0] // 第i天，至多进行k次交易，目前没有(0/1)持有股票
*/
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n)
	// (question) base case
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[0]
			continue
		}
		if i == 1 {
			dp[i][0] = MaxInt(dp[0][0], dp[0][1]+prices[1])
			dp[i][1] = MaxInt(dp[0][0]-prices[1], dp[0][1])
			continue
		}

		// 未持有：前一天未持有（rest）、前一天持有（第i天sell）
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有：前一天持有（rest）、前前一天未持有（第i天buy）
		dp[i][1] = MaxInt(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}

/*
	动态规划，状态压缩(question)
*/
func maxProfit(prices []int) int {
	n := len(prices)
	// base case
	dp_i_0 := 0
	dp_i_1 := -(1 << 32)
	dp_pre_0 := 0 // 代表dp[i-2][0]
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = MaxInt(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = MaxInt(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = temp
	}
	return dp_i_0
}