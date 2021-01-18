/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

	说明:

	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

	思路：动态规划（循环、递归两种方式实现）

	1、状态表示（属性、集合）
		（1）所有s[1~i]和p[1~j]的匹配方案
		（2）bool，是否存在一个合法方案
	2、状态计算
		（1）p[i]!="*",(s[i]==p[j] || p[j]=="." ) && dp[i-1][j-1]
		（2）p[i]=="*",dp[i][j-2] || dp[i-1][j-2] && dp[i][j]

	状态转移方程有点复杂

	"aab"和"c*a*b"测试不通过
*/
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	// 特殊操作
	s = " " + s
	p = " " + p
	// 下表从1开始
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	// s、p都是空的情况
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		// 第二个字符串从1开始
		for j := 1; j <= m; j++ {
			// *不能单独使用
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			// i从1开始
			if i > 0 && p[j] != '*' {
				//
				dp[i][j] = dp[i-1][j-1] && (s[i] == p[i] || p[j] == '.')
			} else if p[j] == '*' {
				/*
					dp[i][j-2],"*"表示0个
					dp[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')，“*”表示一个或多个
				*/
				dp[i][j] = dp[i][j-2] || i > 0 && dp[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[n][m]
}

/*
	递归思路，连这个也不太理解
*/
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	first_match := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
	// 长度大于1的情况
	if len(p) > 1 && p[1] == '*' {
		// p 跳过两个字符，*前面的字符出现0次
		// p 不变，例如 s = aa ，p = a*，第一个 a 匹配，然后 text 的第二个 a 接着和 p 的第一个 a 进行匹配。表示 * 用前一个字符替代。
		return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	}
	// 第一个字符匹配，从第二个字符开始
	return first_match && isMatch(s[1:], p[1:])
}