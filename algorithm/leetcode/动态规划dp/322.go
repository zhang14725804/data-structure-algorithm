/*
	给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
	你可以认为每种硬币的数量是无限的。

	1 <= coins.length <= 12
	1 <= coins[i] <= 231 - 1
	0 <= amount <= 104
*/
func coinChange1(coins []int, amount int) int {
	maxInt := (1 << 32)
	// (question)重叠子问题
	// 输入：[1,2,5] 100，超时
	var dfs func(amount int) int
	// 暴力递归（自顶向下）
	dfs = func(n int) int {
		// base case
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		// 求最小值，所以初始化为正无穷
		res := maxInt
		for _, coin := range coins {
			sub := dfs(n - coin)
			// 子问题无解，跳过
			if sub == -1 {
				continue
			}
			res = MinInt(res, 1+sub)
		}
		if res != maxInt {
			return res
		}
		return -1
	}
	return dfs(amount)
}

func coinChange2(coins []int, amount int) int {
	maxInt := (1 << 32)
	// (question) 通过缓存消除**重叠子问题**
	hash := make(map[int]int)
	var dfs func(amount int) int
	// 暴力递归（自顶向下）
	dfs = func(n int) int {
		if val, ok := hash[n]; ok {
			return val
		}
		// base case
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		// 求最小值，所以初始化为正无穷
		res := maxInt
		for _, coin := range coins {
			sub := dfs(n - coin)
			// 子问题无解，跳过
			if sub == -1 {
				continue
			}
			res = MinInt(res, 1+sub)
		}

		if res != maxInt {
			hash[n] = res
			return res
		}
		hash[n] = -1
		return -1
	}
	return dfs(amount)
}

// 动态规划(question)不懂（😅）
func coinChange(coins []int, amount int) int {
	// ******dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。******
	// 为啥 dp 数组初始化为 amount + 1
	// 因为凑成 amount 金额的硬币数最多只可能等于 amount（全用 1 元面值的硬币），所以初始化为 amount + 1 就相当于初始化为正无穷，便于后续取最小值
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// base case
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = MinInt(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}