/*
	给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

	两个性质：
		（1）左右括号数相等
		（2）任意前缀中，左括号数量>=右括号

	todo：分析挺绕的
*/
func longestValidParentheses(s string) int {
	stk := make([]int, 0)

	res := 0
	for i, start := 0, -1; i < len(s); i++ {
		if s[i] == '(' {
			stk = append(stk, i)
		} else {
			if len(stk) > 0 {
				// stk.pop
				stk = stk[:len(stk)-1]
				if len(stk) > 0 {
					// stk.top
					res = compare(res, i-stk[len(stk)-1], true)
				} else {
					res = compare(res, i-start, true)
				}
			} else {
				start = i
			}
		}
	}
	return res
}
func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}