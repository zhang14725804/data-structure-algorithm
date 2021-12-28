/*
	多次买卖股票，你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
	k=+infinity
*/

/*
	方法2：状态机
	dp[n][k][j]
	如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的，k可以忽略
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		// base case
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 未持有：前一天未持有（rest）、前一天持有（第i天sell）
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持有：前一天持有（rest）、前一天未持有（第i天buy）
		dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/*
	方法2：状态机（优化版），状态压缩
	新状态只和相邻的一个状态有关
*/
func maxProfit(prices []int) int {
	n := len(prices)
	// 未开始，未持有。利润为0
	dp_i_0 := 0
	//  未开始，不可能持有。表示不可能
	dp_i_1 := -(1 << 32)
	for i := 0; i < n; i++ {
		// 未持有
		dp_i_0 = MaxInt(dp_i_0, dp_i_1+prices[i])
		// 持有
		dp_i_1 = MaxInt(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}