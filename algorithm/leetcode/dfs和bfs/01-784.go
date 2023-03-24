
var ans []string

func letterCasePermutation(s string) []string {
	dfs([]rune(s), 0)
	return ans
}

func dfs(s []rune, idx int) {
	// base case
	if idx == len(s) {
		ans = append(ans, string(s))
		return
	}
	// 😅 不变的情况递归一次
	dfs(s, idx+1)
	// 😅 如果当前字符是字母，变大写之后再来一次
	if s[idx] >= 'A' {
		// 😅 转换大小写
		s[idx] ^= 32
		dfs(s, idx+1)
	}
}