/*
	你最多可以完成 两笔 交易。你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	k=2
*/

/*
	方法1：前后缀分解（不懂😅）
*/
func maxProfit1(prices []int) int {
	n := len(prices)

	f := make([]int, n+2) // 这里slice长度为什么n+2
	minp := prices[0]
	// 前缀最大值
	for i := 1; i <= n; i++ {
		f[i] = compare(f[i-1], prices[i-1]-minp, true)
		minp = compare(minp, prices[i-1], false)
	}
	// todo：代码不理解
	res := 0
	maxp := 0
	// 后缀和
	for i := n; i > 0; i-- {
		res = compare(res, maxp-prices[i-1]+f[i-1], true)
		maxp = compare(maxp, prices[i-1], true)
	}

	return res
}

/*
	方法2：动态规划；穷举所有状态(天数i，交易次数k，当前持有状态，三种情况穷举)
*/
func maxProfit(prices []int) int {
	// 初始化
	n, max_k := len(prices), 2
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
			dp[i][k][0] = MaxInt(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = MaxInt(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][max_k][0]
}