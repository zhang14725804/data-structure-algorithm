/*
	给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
	你可以假设 n 不小于 2 且不大于 58。
*/

/*
	动态规划：

		对于的正整数 n，当 n /2 ≥2 时，可以拆分成至少两个正整数的和。
		令 k 是拆分出的第一个正整数，则剩下的部分是 n-k，n-k 可以不继续拆分，
		或者继续拆分成至少两个正整数的和。由于每个正整数对应的最大乘积取决于比它小的正整数对应的最大乘积，
		因此可以使用动态规划求解。

	将 i 拆分成 j 和 i-j 的和，且 i-j 不再拆分成多个正整数，此时的乘积是 j * (i−j)；
	将 i 拆分成 j 和 i-j 的和，且 i-j 继续拆分成多个正整数，此时的乘积是 j * dp[i−j]。

*/
func integerBreak(n int) int {
	// dp[i]：分拆数字i，可以得到的最大乘积为dp[i]。
	dp := make([]int, n+1)
	// dp数组初始化
	dp[2] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// 😅😅😅 递推公式
			dp[i] = MaxInt(dp[i], MaxInt(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

/*
	贪心算法 😅
	每次拆成n个3，如果剩下是4，则保留4，然后相乘
*/
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	if n == 4 {
		return 4
	}
	res := 1
	for n > 4 {
		res *= 3
		n -= 3
	}
	res *= n
	return res
}