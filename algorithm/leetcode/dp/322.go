/*
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
	你可以认为每种硬币的数量是无限的。
*/
func coinChange(coins []int, amount int) int {
	// 输入：[1,2,5] 100，超时
	var dp func(amount int) int
	dp = func(n int) int {
		// base case
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}

		res := (1 << 32)
		for _, coin := range coins {
			sub := dp(n - coin)
			if sub == -1 {
				continue
			}
			res = MinInt(res, 1+sub)
		}
		if res != (1 << 32) {
			return res
		}
		return -1
	}
	return dp(amount)
}