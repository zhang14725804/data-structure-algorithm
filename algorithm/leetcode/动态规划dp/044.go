/*
	给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

	'?' 可以匹配任何单个字符。
	'*' 可以匹配任意字符串（包括空字符串）。
	两个字符串完全匹配才算匹配成功。

	说明:
		s 可能为空，且只包含从 a-z 的小写字母。
		p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
*/

/*
	方法1：动态规划（平平无奇的动态规划）
*/
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	// 用 dp[i][j]dp[i][j] 表示字符串 ss 的前 ii 个字符和模式 pp 的前 jj 个字符是否能匹配
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		// 因为星号才能匹配空字符串，所以只有当模式 p的前 j 个字符均为星号时，dp[0][j]才为真。
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

/*
	方法2：迭代（双指针），不太好理解（question）
*/
func isMatch1(str string, pattern string) bool {
	s, p, match, startIdx := 0, 0, 0, -1
	for s < len(str) {
		if p < len(pattern) && (pattern[p] == '?' || str[s] == pattern[p]) {
			s++
			p++
		} else if p < len(pattern) && pattern[p] == '*' {
			// 碰到 *，假设它匹配空串，并且用 startIdx 记录 * 的位置，记录当前字符串的位置，p 后移
			startIdx = p
			match = s
			p++
		} else if startIdx != -1 {
			// 当前字符不匹配，并且也没有 *，回退；p 回到 * 的下一个位置；match 更新到下一个位置；s 回到更新后的 match
			// 这步代表用 * 匹配了一个字符
			p = startIdx + 1
			match++
			s = match
		} else {
			return false
		}
	}
	for p < len(pattern) && pattern[p] == '*' {
		p++
	}
	return p == len(pattern)
}