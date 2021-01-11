/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
	有效字符串需满足：
		左括号必须用相同类型的右括号闭合。
		左括号必须以正确的顺序闭合。
		注意空字符串可被认为是有效字符串。

	思路：遍历字符串，存入栈中，和栈顶元素进行比较。思路经典
*/
func isValid(s string) bool {
	stack := make([]byte, 0)
	// for循环字符串 byte类型；range循环字符串rune类型
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && leftOf(c) == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func leftOf(c byte) byte {
	if c == '}' {
		return '{'
	}
	if c == ']' {
		return '['
	}
	return '('
}