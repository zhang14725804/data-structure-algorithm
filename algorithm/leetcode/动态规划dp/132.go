/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回符合要求的最少分割次数。
*/

/*
	question
	😅😅😅
	难，动态规划的思路
	todo：思路不懂，代码也有问题
*/
func minCut(s string) int {
	s = " " + s
	n := len(s)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		f[i] = INT_MAX
	}
	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	// 考虑所有长度的子串（最小长度从1开始）
	for len := 1; len <= n; len++ {
		// 从每个下标开始
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			// i + 1 <= j - 1 并且 j:=i+len-1，所有len<3
			dp[i][j] = s[i] == s[j] && (len < 3 || dp[i+1][j-1])
		}
	}

	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if dp[j][i] {
				f[i] = MinInt(f[i], f[j-1]+1)
			}
		}
	}
	return f[n] - 1
}

/*
	state: f[i] "前i"个字符组成的子字符串需要最少几次cut(个数-1为索引)
    function: f[i] = MIN{f[j]+1}, j < i && [j+1 ~ i]这一段是一个回文串
    intialize: f[i] = i - 1 (f[0] = -1)
    answer: f[s.length()]
*/
func minCut(s string) int {
	sLen := len(s)
	if sLen == 0 || sLen == 1 {
		return 0
	}

	dp := make([]int, sLen+1)
	dp[0] = -1
	dp[1] = 0
	for i := 1; i <= sLen; i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				dp[i] = MinInt(dp[i], dp[j]+1)
			}
		}
	}
	return dp[sLen]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}