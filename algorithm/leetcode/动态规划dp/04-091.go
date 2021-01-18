package leetcode

/*
	一条包含字母 A-Z 的消息通过以下方式进行了编码：

	'A' -> 1
	'B' -> 2
	...
	'Z' -> 26
	给定一个只包含数字的非空字符串，请计算解码方法的总数。

	动态规划思路：
	（1）状态表示。集合（所有由前i个数字解码得到的字符串），属性：数量
	（2）状态计算。最后一个字母是1位数f[i-1]，最后一个字母是两位数f[i-2]

	由递归再去想动态规划，就好理解了
*/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	// 声明数组
	dp := make([]int, n+1)
	// 0个数字，解码空字符串
	dp[0] = 1
	// 第一个字符串
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		// 最后一个字母是一位数
		// last, _ := strconv.Atoi(s[i-1 : i])
		last := s[i-1] - '0'
		if last >= 1 && last <= 9 {
			dp[i] += dp[i-1]
		}
		// 最后一个字母是两位数
		// last, _ = strconv.Atoi(s[i-2 : i])
		last = (s[i-2]-'0')*10 + (s[i-1] - '0')
		if last >= 10 && last <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

/*
	思路2：递归
*/
func numDecodings(s string) int {
	return dfs(s, 0)
}
func dfs(s string, i int) int {
	if i == len(s) {
		return 1
	}
	if s[i] == '0' {
		return 0
	}
	ans1 := dfs(s, i+1)
	ans2 := 0
	// 这里为什么不需要判断 >=10的情况
	if i < len(s)-1 {
		if (s[i]-'0')*10+(s[i+1]-'0') <= 26 {
			ans2 = dfs(s, i+2)
		}
	}
	return ans1 + ans2
}
