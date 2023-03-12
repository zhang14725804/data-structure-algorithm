/*
	1. 从前向后遍历，【stack存储当前数字】，默认符号是“+”号，默认num=0
	2. 遇到数字转数字
	3. 遇到非数字非空格（也就是遇到加减乘除）或者已经遍历完时，根据当前sign把当前数字入栈，更新num和sign
*/

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	return helper(s)
}

func helper(s string) int {
	stack := make([]int, 0)
	num := 0
	sign := '+'

	// 从前向后遍历
	for len(s) > 0 {
		c := s[0]
		s = s[1:]

		// 是数字
		if isdigit(c) {
			num = num*10 + int(c-'0')
		}

		// 非数字并且非空格（遇到符号），或者字符串长度为0
		if (!isdigit(c) && c != ' ') || len(s) == 0 {
			switch sign {
			case '-':
				stack = append(stack, -num)
			case '+':
				stack = append(stack, num)
			case '*':
				pre := stack[len(stack)-1]   // top
				stack = stack[:len(stack)-1] // pop
				stack = append(stack, num*pre)
			case '/':
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pre/num)
			}

			// 😅
			sign = rune(c)
			num = 0
		}
	}

	return sum(stack)
}

func sum(stack []int) int {
	res := 0
	for i := 0; i < len(stack); i++ {
		res += stack[i]
	}
	return res
}

func isdigit(char byte) bool {
	return char >= '0' && char <= '9'
}
