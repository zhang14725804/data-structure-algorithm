/*
	给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

	括号序列：
		（1）每个左括号唯一对应一个右括号
		（2）一个括号序列合法  <==> （若：左括号记为1，右括号记为-1）所有前缀和>=0,且总和==0
		也就是左括号数量大于等于右括号，且左右括号数量相等
	(()()) ==== 1 1 -1 1 -1 1

*/

/*
	方法1：正着判断一次，反转之后在判断一次
*/
func longestValidParentheses1(s1 string) int {
	// 转为rune数组
	s := []rune(s1)
	// 正着做一次
	res := work(s)
	// (question)反转，(没想到反转😅)，变换左右括号之后再做一次
	s = reverse(s)
	// for _, c := range s {
	// 	c ^= 1 // 需要修改数组元素，这样修改是错的😅
	// }
	// 左括号换成右括号（因为他们的ascii值只是最后一位相差1，所以直接用^=操作）
	for i, _ := range s {
		s[i] ^= 1
	}
	return MaxInt(res, work(s))
}

/*
	分为三种情况(根据括号序列的性质)
	cnt<0，当前子段不合法，start=i,cnt=0
	cnt>0，继续
	cnt==0，[start,i]是合法的序列
*/
func work(s []rune) int {
	res := 0
	for i, start, cnt := 0, 0, 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
			if cnt < 0 {
				start = i + 1
				cnt = 0
			} else if cnt == 0 {
				res = MaxInt(res, i-start+1)
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

/*
	利用栈
*/
func longestValidParentheses2(s string) int {
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

/*
	动态规划(question)
	dp [ i ] 代表以下标 i 结尾的合法序列的最长长度
*/