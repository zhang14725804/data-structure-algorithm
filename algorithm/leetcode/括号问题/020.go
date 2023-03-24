/*
	1. 遇到左括号直接入栈
	2. 遇到右括号，判断是否和当前栈顶括号匹配，若匹配则出栈，进行下一轮
	3. 最后判断stack==0
*/
func isValid(s string) bool {
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