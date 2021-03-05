/*
	买卖股票的最佳时机含手续费
	k = +infinity with fee
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		// base case
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 第i天未持有：前一天未持有（rest）、前一天持有（第i天sell,😅减去手续费）
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		// 第i天持有：前一天持有（rest）、前一天未持有（第i天buy）
		dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	fmt.Println(dp)
	return dp[n-1][0]
}

/*
	动态规划，状态压缩(question)

*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	// base case
	dp_i_0 := 0
	dp_i_1 := -prices[0]
	for i := 0; i < n; i++ {
		dp_i_0 = MaxInt(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = MaxInt(dp_i_1, dp_i_0-prices[i]-fee)
	}
	return dp_i_0
}