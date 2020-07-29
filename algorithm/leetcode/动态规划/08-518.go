/*
	给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

	动态规划(背包问题)
*/
func change(m int, coins []int) int {
	/*
		集合：所有由前i种硬币凑出的总钱数为j的凑发
		属性：所有凑发的数量

		状态计算：f[i][j]，

		todo：不懂

		f[i][j] = f[i-1,j]+f[i-1,j-c],f[i-1,j-2c]+...+f[i-1,kc]
		优化后(从二维到一维)：
		f[i][j] = f[i-1,j]+f[i,j-c]
		f[j] = f[j]+f[j-c]
	*/
	n := len(coins)
	f := make([]int, m+1)

	f[0] = 1
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= m; j++ {
			f[j] += f[j-coins[i]]
		}
	}
	return f[m]
}
