/*
	给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

	两个性质：
		（1）左右括号数相等
		（2）任意前缀中，左括号数量>=右括号

	todo：分析挺绕的
*/
func longestValidParentheses(s string) int {
	// 用slice模拟栈
	stk := make([]int, 0)
	res := 0
	// start，记录每个分段的前一个位置
	for i, start := 0, -1; i < len(s); i++ {
		// 当前这个字符是左括号
		if s[i] == '(' {
			stk = append(stk, i)
		} else {
			// 当前是右括号
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
				// 当前右括号是分界点
				start = i
			}
		}
	}
	return res
}
