/*
	有台奇怪的打印机有以下两个特殊要求：

	打印机每次只能打印同一个字符序列。
	每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
	给定一个只包含小写英文字母的字符串，你的任务是计算这个打印机打印它需要的最少次数。

	todo:不懂
*/
func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	//
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 0; j+i-1 < n; j++ {
			r := j + i - 1
			dp[j][r] = dp[j+1][r] + 1
			for k := j + 1; k <= r; k++ {
				if s[k] == s[j] {
					dp[j][r] = compare(dp[j][r], dp[j][k-1]+dp[k+1][r], false)
				}
			}
		}
	}
	return dp[0][n-1]
}