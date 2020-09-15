/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。
*/

// 方法1：分治
var s string

func partition(_s string) [][]string {
	s = _s
	return dfs(0)
}

func dfs(u int) [][]string {
	ans := make([][]string, 0)
	// 递归出口
	if u == len(s) {
		path = make([]string, 0)
		ans = append(ans, path)
		return ans
	}
	// 遍历每个切割的位置
	for i := u; i < len(s); i++ {
		if check(s[u : i+1]) {
			left := s[u : i+1]
			for _, c := range dfs(i + 1) {
				c = append([]string{left}, c...)
				ans = append(ans, c)
			}
		}
	}
	return ans
}

// 当前区间是否是回文串
func check(now string) bool {
	if len(now) == 0 {
		return false
	}
	for i, j := 0, len(now)-1; i < j; i, j = i+1, j-1 {
		if now[i] != now[j] {
			return false
		}
	}
	return true
}

// 方法12：分治优化。用动态规划的方法，把所有字符是否是回文串提前存起来
var dp [][]bool
var s string

func partition(_s string) [][]string {
	s = _s
	n := len(s)
	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for len := 1; len <= n; len++ {
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			// i + 1 <= j - 1 && j = i + len - 1。所以：len < 3
			dp[i][j] = s[i] == s[j] && (len < 3 || dp[i+1][j-1])
		}
	}

	return dfs(0)
}

func dfs(u int) [][]string {
	ans := make([][]string, 0)
	if u == len(s) {
		path := make([]string, 0)
		ans = append(ans, path)
		return ans
	}

	for i := u; i < len(s); i++ {
		if dp[u][i] {
			left := s[u : i+1]
			for _, l := range dfs(i + 1) {
				l = append([]string{left}, l...)
				ans = append(ans, l)
			}
		}
	}
	return ans
}

// 方法3：回溯
var ans [][]string
var dp [][]bool
var s string

func partition(_s string) [][]string {
	s = _s
	ans = make([][]string, 0)

	n := len(s)
	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for len := 1; len <= n; len++ {
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			// i + 1 <= j - 1 && j = i + len - 1。所以：len < 3
			dp[i][j] = s[i] == s[j] && (len < 3 || dp[i+1][j-1])
		}
	}

	backtrack(0, make([]string, 0))
	return ans
}
func backtrack(u int, temp []string) {
	if u == len(s) {
		ans = append(ans, temp)
	}
	for i := u; i < len(s); i++ {
		if dp[u][i] {
			left := s[u : i+1]
			temp = append(temp, left)
			backtrack(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
}