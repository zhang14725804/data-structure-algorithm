/*
	给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
	你可以假设 n 不小于 2 且不大于 58。
*/

/*
	动态规划(question),平平无奇的动态规划：
		对于的正整数 nn，当 n \ge 2n≥2 时，可以拆分成至少两个正整数的和。
		令 kk 是拆分出的第一个正整数，则剩下的部分是 n-kn−k，n-kn−k 可以不继续拆分，
		或者继续拆分成至少两个正整数的和。由于每个正整数对应的最大乘积取决于比它小的正整数对应的最大乘积，
		因此可以使用动态规划求解。

	将 i 拆分成 j 和 i-j 的和，且 i-j 不再拆分成多个正整数，此时的乘积是 j * (i−j)；
	将 i 拆分成 j 和 i-j 的和，且 i-j 继续拆分成多个正整数，此时的乘积是 j * dp[i−j]。

*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = MaxInt(curMax, MaxInt(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}