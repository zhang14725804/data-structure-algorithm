/*
	方法1：
		首先，判断合法字符的长度：
		1. cnt 记录左右括号数量，左括号+1，右括号-1
		2. start 记录起始位置
		3. 遍历字符串
		分为三种情况(根据括号序列的性质)
			cnt<0，当前子段不合法，start=i,cnt=0
			cnt>0，继续
			cnt==0，[start,i]是合法的序列
		然后，正着判断一次，反转之后在判断一次
*/
func longestValidParentheses1(s1 string) int {
	// 转为rune数组
	s := []rune(s1)
	// 正着做一次
	res := work(s)
	// 反转字符串
	s = reverse(s)
	// 左括号换成右括号（因为他们的ascii值只是最后一位相差1，所以直接用^=操作）
	// c ^= 1 // 需要修改数组元素，这样修改是错的😅
	for i, _ := range s {
		s[i] ^= 1
	}
	// (question)反转，(没想到反转😅)，变换左右括号之后再做一次
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
	// 记录左右括号数量，左括号+1，右括号-1
	cnt := 0
	// 记录起始位置
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 左括号
			cnt++
		} else {
			// 右括号
			cnt--
			// 不合法序列，跳过
			if cnt < 0 {
				start = i + 1
				cnt = 0
			} else if cnt == 0 {
				// 合法序列
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
	方法2:利用栈 TODO 😅
		从左到右扫描字符串，栈顶保存当前扫描的时候，合法序列前的一个位置位置下标是多少
			1. 扫描到左括号，就将当前位置入栈
			2. 扫描到右括号，就将栈顶出栈（代表栈顶的左括号匹配到了右括号），然后分两种情况
				栈不空，那么就用当前的位置减去栈顶的存的位置，然后就得到当前合法序列的长度，然后更新一下最长长度。
				栈是空的，说明之前没有与之匹配的左括号，那么就将当前的位置入栈。
*/
func longestValidParentheses2(s string) int {
	// 用slice模拟栈
	stk := make([]int, 0)
	res := 0
	// start，记录每个分段的前一个位置
	start := -1
	for i := 0; i < len(s); i++ {
		// 当前这个字符是左括号
		if s[i] == '(' {
			stk = append(stk, i)
		} else {
			// 当前是右括号
			if len(stk) > 0 {
				// 出栈
				stk = stk[:len(stk)-1]
				if len(stk) > 0 {
					// stk.top
					res = Max(res, i-stk[len(stk)-1])
				} else {
					res = Max(res, i-start)
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
	TODO 😅
*/