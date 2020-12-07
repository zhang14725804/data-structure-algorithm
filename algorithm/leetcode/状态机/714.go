/*
	买卖股票的最佳时机含手续费
	k = +infinity with fee
*/
func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	// (question)；状态转移方程有问题，fee该在那个地方处理
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}

/*
	动态规划，状态压缩
*/
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := -(1 << 32)
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = MaxInt(temp, dp_i_1+prices[i])
		dp_i_1 = MaxInt(dp_i_1, temp-prices[i]-fee)
	}
	return dp_i_0
}