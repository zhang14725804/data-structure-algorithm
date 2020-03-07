/*
	32. Longest Valid Parentheses

	括号序列性质：
	每个左括号对应右括号确定？
	括号序列合法(-1,1)，所有前缀和>=0，而且总和=0

	(()()) ==== 1 1 -1 1 -1 1

*/
func longestValidParentheses(s1 string) int {
	s := []rune(s1)
	// 正着做一次
	res := work(s)
	// 反转，变换左括号右括号之后再做一次
	s = reverse(s)
	// 左括号换成右括号（因为他们的ascii相差1，所以直接用^=操作）
	for _, c := range s {
		c ^= 1
	}
	return max(res, work(s))
}

/*
	分为三种情况(根据括号序列的性质)
	cnt<0
	cnt>0
	cnt==0

	todos::这里有问题（断点调试是个问题）
*/
func work(s []rune) int {
	res := 0
	for i, start, cnt := 0, 0, 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt += 1
		} else {
			cnt -= 1
			if cnt < 0 {
				start = i + 1
				cnt = 0
			} else if cnt == 0 {
				res = max(res, i-start+1)
			}
		}
	}
	return res
}

// 反转字符串
func reverse(a []rune) []rune {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
