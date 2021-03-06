/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

	说明:

	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
*/
/*
	思路：动态规划（循环、递归两种方式实现）

	1、状态表示（属性、集合）
		（1）所有s[1~i]和p[1~j]的匹配方案
		（2）bool，是否存在一个合法方案
	2、状态计算
		（1）p[i]!="*"：(s[i]==p[j] || p[j]=="." ) && dp[i-1][j-1]
		（2）p[i]=="*"：dp[i][j-2] || dp[i-1][j-2] && dp[i][j]

	TODO: "aab"和"c*a*b"测试不通过
*/
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 特殊操作，避规边界条件判断
	s = " " + s
	p = " " + p
	// 下表从1开始
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	// base case: s、p都是空的情况
	dp[0][0] = true
	// i,j从1开始
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// *不能单独使用
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			// s[i]==p[j]或者p[j]=='.'
			firstMatch := s[i] == p[j] || p[j] == '.'
			if p[j] != '*' {
				dp[i][j] = dp[i-1][j-1] && firstMatch
			} else if p[j] == '*' {
				/*
					(1) dp[i][j-2]："*"表示0个
					(2) dp[i-1][j] && firstMatch：“*”表示一个或多个
				*/
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.'))
			}
		}
	}
	return dp[n][m]
}