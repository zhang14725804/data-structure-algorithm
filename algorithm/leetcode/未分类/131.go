/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回 s 所有可能的分割方案。

	思路：s[i,j]是回文串，那么s[i+1,j-1]也是回文串

*/

/*
	思路1：分治
*/
func partition(s string) [][]string {
	return dfs(s, 0)
}

func dfs(s string, u int) [][]string {
	// 递归的出口，返回空的二维数组（todo：为什么对结果不影响呢）
	if u == len(s) {
		path := make([]string, 0)
		ans := make([][]string, 0)
		ans = append(ans, path)
		return ans
	}

	ans := make([][]string, 0)
	// 当前切割后是回文串才考虑
	for i := u; i < len(s); i++ {
		if isPalindrome(s[u : i+1]) {
			left := s[u : i+1]
			// 遍历后边字符串的所有结果，将当前的字符串加到头部
			for _, l := range dfs(s, i+1) {
				// slice头部添加元素
				l = append([]string{left}, l...)
				ans = append(ans, l)
			}
		}
	}
	return ans
}

// 判断字符串是否是回文串
func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/*
	思路2：分治优化，去除重复的判断是否是回文数
*/
var dp [][]bool

func partition(s string) [][]string {
	n := len(s)
	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// 这里是重点
	// 考虑所有长度的子串（最小长度从1开始）
	for len := 1; len <= n; len++ {
		// 从每个下标开始
		for i := 0; i <= n-len; i++ {
			j := i + len - 1
			// i + 1 <= j - 1 并且 j:=i+len-1，所有len<3
			dp[i][j] = s[i] == s[j] && (len < 3 || dp[i+1][j-1])
		}
	}

	return dfs(s, 0)
}

func dfs(s string, u int) [][]string {
	ans := make([][]string, 0)
	if u == len(s) {
		path := make([]string, 0)
		ans = append(ans, path)
		return ans
	}

	for i := u; i < len(s); i++ {
		if dp[u][i] {
			left := s[u : i+1]
			for _, l := range dfs(s, i+1) {
				l = append([]string{left}, l...)
				ans = append(ans, l)
			}
		}
	}
	return ans
}

/*
	思路3：回溯
*/
var dp [][]bool
var ans [][]string

func partition(s string) [][]string {
	n := len(s)
	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
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

	dfs(s, 0, make([]string, 0))
	return ans
}
func dfs(s string, u int, path []string) {
	if u == len(s) {
		ans = append(ans, path)
	}
	//
	for i := u; i < len(s); i++ {
		if dp[u][i] {
			left := s[u : i+1]
			path = append(path, left)
			dfs(s, i+1, path)
			path = path[:len(path)-1]
		}
	}
}