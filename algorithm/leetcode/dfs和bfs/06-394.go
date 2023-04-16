/*
	栈实现 😅😅😅
	1. 是数字，解析数字（连续的多个位数）并进栈
	2. 是字母或者左括号，直接进栈
	3. 右括号，出栈，直到左括号为止
	4. 😅 出栈序列反转后拼接成一个字符串
	5. 删除栈顶左括号，取出栈顶数字，repeat之后入栈
	6. stk 转字符串返回
*/

func decodeString(s string) string {
	stk := []string{}
	ptr := 0

	for ptr < len(s) {
		cur := s[ptr]
		// 是数字，解析数字（连续的多个位数）并进栈
		if cur >= '0' && cur <= '9' {
			dig := getDigits(s, &ptr)
			stk = append(stk, dig)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			// 是字母或者左括号，直接进栈
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			// 右括号，出栈，直到左括号为止
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}

			// 😅 出栈序列反转后拼接成一个字符串
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}

			// 删除栈顶左括号
			stk = stk[:len(stk)-1]
			// 取出栈顶数字
			repeats, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			// repeat之后入栈
			t := strings.Repeat(getString(sub), repeats)
			stk = append(stk, t)

		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}
func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}