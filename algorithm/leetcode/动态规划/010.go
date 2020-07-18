/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

	思路：动态规划（循环、递归两种方式实现）

	1、状态表示（属性、集合）
		（1）所有s[1~i]和p[1~j]的匹配方案
		（2）bool，是否存在一个合法方案
	2、状态计算
		（1）p[i]!="*",(s[i]==p[j] || p[j]=="." ) && dp[i-1][j-1]
		（2）p[i]=="*",dp[i][j-2] || dp[i-1][j-2] && dp[i][j]

	状态转移方程有点复杂
*/
func isMatch(s string, p string) bool {
	
}