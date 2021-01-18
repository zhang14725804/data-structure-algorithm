/*
	给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

	一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
	例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

	若这两个字符串没有公共子序列，则返回 0。
*/

/*
	方法1：递归（自顶向下），重叠子问题

	todo：加cache缓存
*/
func longestCommonSubsequence(s1 string, s2 string) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// base case 空串
		if i == -1 || j == -1 {
			return 0
		}
		// 两种情况
		if s1[i] == s2[j] {
			return dfs(i-1, j-1) + 1
		} else {
			return MaxInt(dfs(i-1, j), dfs(i, j-1))
		}
	}
	return dfs(len(s1)-1, len(s2)-1)
}

/*
	方法2：动态规划(自底向上)
	todo：状态压缩
*/
func longestCommonSubsequence(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			// base case
			dp[i][j] = 0
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = MaxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n][m]
}