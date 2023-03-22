/*
	栈实现
	从左到右遍历布尔表达式，对于每种类型的字符，执行相应的操作：
		1. 如果当前字符是逗号，则跳过该字符；
		2. 如果当前字符是除了逗号和右括号以外的任意字符，则将该字符添加到栈内；
		3. 如果当前字符是右括号，则一个表达式遍历结束，需要解析该表达式的值，并将结果添加到栈内
*/
func parseBoolExpr(expression string) bool {
	stack := []rune{}
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stack = append(stack, c)
			continue
		}

		// 统计t和f的数量
		tc := 0
		fc := 0
		for stack[len(stack)-1] != '(' {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if val == 't' {
				tc++
			} else {
				fc++
			}
		}
		// 左括号出栈
		stack = stack[:len(stack)-1]
		// 运算符出栈
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 处理一个括号中的内容，并入栈
		c := 't'

		// 😅 😅 😅 这里是难点
		switch op {
		case '!':
			// 逻辑非
			if fc != 1 {
				c = 'f'
			}
			stack = append(stack, c)
		case '&':
			// 逻辑与
			if fc != 0 {
				c = 'f'
			}
			stack = append(stack, c)
		case '|':
			// 逻辑或
			if tc == 0 {
				c = 'f'
			}
			stack = append(stack, c)
		}
	}
	return stack[len(stack)-1] == 't'
}