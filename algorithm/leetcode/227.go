/*
	实现一个基本的计算器来计算一个简单的字符串表达式的值。
	字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

	todo：三种方法😅，难
*/

func calculate(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	stack := &Stack{}
	num := 0
	sign := '+'

	// todo：代码有问题
	for i := 0; i < l; i++ {
		fmt.Println(string(s[i]), unicode.IsDigit(rune(s[i])))
		if unicode.IsDigit(rune(s[i])) {
			num = num*10 + int(s[i]-'0')
		} else if !unicode.IsDigit(rune(s[i])) && s[i] != ' ' || i == l-1 {
			fmt.Println(num, string(sign))
			if sign == '-' {
				stack.push(-num)
			}
			if sign == '+' {
				stack.push(num)
			}
			if sign == '*' {
				stack.push(stack.top() * num)
				stack.pop()
			}
			if sign == '/' {
				stack.push(stack.top() / num)
				stack.pop()
			}
			sign = rune(s[i])
			num = 0
		}
	}
	fmt.Println(stack.x)
	res := 0
	for i := 0; i < len(stack.x); i++ {
		res += stack.x[i]
	}
	return res
}