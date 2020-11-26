/*
	给定两个单词 word1 和 word2，找到使得 word1 和 word2 相同所需的最小步数，每步可以删除任意一个字符串中的一个字符。
*/

/*
	方法1：递归（自顶向下）
	超时，重复子序列问题
*/
func minDistance1(s1 string, s2 string) int {
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// base case 空串
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		// 两种情况
		if s1[i] == s2[j] {
			return dfs(i-1, j-1)
		} else {
			return MinInt(dfs(i-1, j), dfs(i, j-1)) + 1
		}
	}
	return dfs(len(s1)-1, len(s2)-1)
}

/*
	方法2：动态规划（自底向上）
*/
func minDistance2(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	// 最长公共子序列问题
	lcs := longestCommonSubsequence(s1, s2)
	return n - lcs + m - lcs
}

/*
	动态规划（自底向上）
*/
func minDistance(s1 string, s2 string) int {
	n, m := len(s1), len(s2)
	if n == 0 && m == 0 {
		return 0
	}
	// dp[i][j] 表示要使 word1 的前 i 个字符和 word2 的前 j 个字符匹配所需要的最小步数
	dp := make([][]int, n+1)
	// base case
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
		for j := 0; j <= m; j++ {
			dp[0][j] = j
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 两种情况
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = MinInt(dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[n][m]
}