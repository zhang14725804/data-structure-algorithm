/*
	思路：遍历字符串，存入【栈】中，和栈顶元素进行比较。
	0102 懵逼状态
*/
func isValid1(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		// （1）左括号直接入栈
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			// （2）右括号判断，是否和栈顶括号匹配。
			// 若匹配出栈继续判断下一个；不匹配直接返回false
			if len(stack) > 0 && leftOff(c) == stack[len(stack)-1] {
				// 出栈
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func leftOff(c byte) byte {
	if c == '}' {
		return '{'
	}
	if c == ']' {
		return '['
	}
	return '('
}

/*
   1. 左括号入栈
   2. 遇到右括号，检查是否匹配，若匹配则出栈
   3. 最后检查栈长度是否为0
*/
func isValid2(s string) bool {
	sb := []byte(s)
	stack := make([]byte, 0)
	for i := 0; i < len(sb); i++ {
		switch sb[i] {
		case '(', '[', '{':
			stack = append(stack, sb[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			tp := stack[len(stack)-1]
			if (tp == '(' && sb[i] == ')') || (tp == '[' && sb[i] == ']') || (tp == '{' && sb[i] == '}') {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, sb[i])
			}
		}
	}
	return len(stack) == 0
}