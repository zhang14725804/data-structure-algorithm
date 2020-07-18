/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

	说明:

	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

	todo:动态规划的分析
	"aab"和"c*a*b"测试不通过

	思路：动态规划
*/
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 特殊操作
	s = " " + s
	p = " " + p

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 0; i <= n; i++ {
		//
		for j := 1; j <= m; j++ {
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i > 0 && p[j] != '*' {
				dp[i][j] = dp[i-1][j-1] && (s[i] == p[i] || p[j] == '.')
			} else if p[j] == '*' {
				dp[i][j] = dp[i][j-2] || i > 0 && dp[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[n][m]
}