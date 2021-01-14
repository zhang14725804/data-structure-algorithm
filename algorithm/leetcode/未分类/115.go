/*
	给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
	一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
	（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
	题目数据保证答案符合 32 位带符号整数范围。

	思路：动态规划
		（1）状态表示；集合（s[1~i]）和t[1~j]相等的子序列，属性（数量）
		（2）状态计算；选或者不选s[i]
*/
func numDistinct(s string, t string) int {
	n,m:=len(s),len(t)
	s=" "+s
	t=" "+t
	// 
	dp:=make([][]int,n+1)
	for i := 0; i <= n; i++ {
		dp[i]=make([]int,m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	// 为什么要从1开始
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// todo：状态计算，不懂
			dp[i][j] = dp[i-1][j]
			if s[i]==t[j] {
				dp[i][j] = dp[i][j]+dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}