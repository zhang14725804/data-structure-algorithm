/*

 */
func minCostClimbingStairs(cost []int) int {
	cLen := len(cost)
	// dp[i]的定义：到达第i个台阶所花费的最少体力为dp[i]
	dp := make([]int, cLen)
	// dp数组如何初始化
	dp[0] = cost[0]
	dp[1] = cost[1]
	// 确定遍历顺序
	for i := 2; i < cLen; i++ {
		// 递推公式
		dp[i] = MinInt(dp[i-1], dp[i-2]) + cost[i]
	}
	// 注意最后一步可以理解为不用花费，所以取倒数第一步，第二步的最少值
	return MinInt(dp[cLen-1], dp[cLen-2])
}

/*
	状态压缩 😅😅😅
*/