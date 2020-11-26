/*
	给定一个字符串 s ，找到其中最长的回文子序列，并返回该序列的长度。可以假设 s 的最大长度为 1000 。
	给定数据范围有什么作用(question)
*/
/*
	方法1：动态规划

		状态：
		f[i][j] 表示 s 的第 i 个字符到第 j 个字符组成的子串中，最长的回文序列长度是多少。

		转移方程：
			如果 s 的第 i 个字符和第 j 个字符相同的话
			f[i][j] = f[i + 1][j - 1] + 2
			如果 s 的第 i 个字符和第 j 个字符不同的话
			f[i][j] = max(f[i + 1][j], f[i][j - 1])

		然后注意😅遍历方向(question)：
		为了保证每次计算 dp[i][j]，左下右方向的位置已经被计算出来，只能斜着遍历或者反着遍历(从下到上)
		遍历顺序，i 从最后一个字符开始往前遍历，j 从 i + 1 开始往后遍历，这样可以保证每个子问题都已经算好了。

		初始化：
		f[i][i] = 1

		结果：
		f[0][n - 1]
*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		// 单个字符的最长回文序列是 1。如果只有一个字符，显然最长回文子序列长度是 1，也就是 dp[i][j] = 1 (i == j)
		dp[i][i] = 1
		// i一直是小于j的，不存在i大于j的情况
		for j := i + 1; j < n; j++ {
			// 状态转移方程
			// 对 dp[i][j] 的更新，其实只依赖于 dp[i+1][j-1], dp[i][j-1], dp[i+1][j] 这三个状态
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = MaxInt(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

/*
	动态规划（进行状态压缩）,难以理解
	(question)状态压缩
*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	// base case：一维 dp 数组全部初始化为 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		prev := 0
		// i一直是小于j的，不存在i大于j的情况
		for j := i + 1; j < n; j++ {
			temp := dp[j]
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = MaxInt(dp[j], dp[j-1])
			}
			prev = temp
		}
	}
	return dp[n-1]
}