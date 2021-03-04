/*
	一次买卖股票
	k=1
*/

/*
	方法1：暴力做法
*/
func maxProfit1(prices []int) int {
	res := 0
	// 两次循环直接结算最大值
	for i := 1; i < len(prices); i++ {
		for j := 0; j < i; j++ {
			res = MaxInt(prices[i]-prices[j], res)
		}
	}
	return res
}

/*
	方法1：暴力做法优化
*/
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	minp := prices[0]
	// 一边扫描一边记录最小值，一遍扫描搞定
	for i := 0; i < len(prices); i++ {
		res = MaxInt(res, prices[i]-minp)
		minp = MaxInt(minp, prices[i])
	}
	return res
}

/*
	方法2：状态机（动态规划）
	dp[i][k][0] // 第i天，至多进行k次交易，目前没有(0/1)持有股票

	buy（买）、sell（卖）、rest（无操作）

	k=0,dp[i-1][0][0]=0
	k=1不会改变，
	k可以省略
*/
func maxProfit3(prices []int) int {
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
		// base case
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 第i天未持有：前一天未持有（rest）、前一天持有（第i天sell）
		dp[i][0] = MaxInt(dp[i-1][0], dp[i-1][1]+prices[i])
		// 第i天持有：前一天持有（rest）、前一天未持有（第i天buy）
		dp[i][1] = MaxInt(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

/*
	方法2：状态机（动态规划），优化版（TODO：状态压缩）
	新状态只和相邻的一个状态有关
*/
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 未开始，未持有。利润为0
	dp_i_0 := 0
	// 未开始，不可能持有。表示不可能
	dp_i_1 := -(1 << 32)
	for i := 0; i < n; i++ {
		// 未持有
		dp_i_0 = MaxInt(dp_i_0, dp_i_1+prices[i])
		// 持有
		dp_i_1 = MaxInt(dp_i_1, -prices[i])
	}
	return dp_i_0
}