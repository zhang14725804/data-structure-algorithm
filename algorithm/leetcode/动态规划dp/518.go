/*
	给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
*/

/*
	解题思路：

	第一步要明确两点，「状态」和「选择」
	状态有两个，就是「背包的容量」和「可选择的物品」
	选择就是「装进背包」或者「不装进背包」

	第二步要明确dp数组的定义
	只使用coins中的前i个硬币的面值，若想凑出金额j，有dp[i][j]种凑法
	base case 为dp[0][..] = 0， dp[..][0] = 1。因为如果不使用任何硬币面值，就无法凑出任何金额；如果凑出的目标金额为 0，那么“无为而治”就是唯一的一种凑法。

	第三步，根据「选择」，思考状态转移的逻辑
	不选第i个物品：dp[i][j] = dp[i-1][j]
	选第i个物品：dp[i][j] = dp[i][j-coins[i-1]]
*/
func change1(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		// base case 凑出的目标金额为 0
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			// 剩余金额充足
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}

/*
	状态压缩( question )
*/
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	// base case
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
