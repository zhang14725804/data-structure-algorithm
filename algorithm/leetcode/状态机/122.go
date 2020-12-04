/*
	多次买卖股票，你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）
	k=+infinity
*/

/*
	方法1(思路妙啊)：每次交易分解成每天都交易，把所有结果为正数的加起来
*/
func maxProfit1(prices []int) int {
	res := 0
	for i := 0; i+1 < len(prices); i++ {
		res += MaxInt(0, prices[i+1]-prices[i])
	}
	return res
}

/*
	方法2：状态机
	如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的
	(question)不理解😅
*/
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 初始化操作
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = MaxInt(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

/*
	方法2：状态机（优化版），状态压缩
*/
func maxProfit(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := -(1 << 32)
	for i := 0; i < n; i++ {
		temp := dp_i_0
		dp_i_0 = MaxInt(temp, dp_i_1+prices[i])
		dp_i_1 = MaxInt(dp_i_1, temp-prices[i])
	}
	return dp_i_0
}