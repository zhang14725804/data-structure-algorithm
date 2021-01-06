/*
	实现一个基本的计算器来计算一个简单的字符串表达式的值。
	字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

*/

/*
	question 【处理复杂问题的思路】
	（1）字符串转数字
	（2）处理只包含加减法的算式
	（3）处理包含加减乘除四则运算的算式
	（4）处理空格字符
	（5）处理包含括号的算式
*/
func calculate(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	return dfs(s)
}

func dfs(s string) int {
	stack := make([]int, 0)
	num := 0
	sign := '+'
	for len(s) > 0 {
		c := s[0]
		s = s[1:] // pop(0)
		if isdigit(c) {
			num = num*10 + int(c-'0')
		}
		// 遇到左括号，递归
		if c == '(' {
			num = dfs(s)
		}
		if (!isdigit(c) && c != ' ') || len(s) == 0 {
			switch sign {
			case '-':
				fmt.Println("减")
				stack = append(stack, -num)
			case '+':
				fmt.Println("加")
				stack = append(stack, num)
			case '*':
				fmt.Println("乘")
				pre := stack[len(stack)-1]   // top
				stack = stack[:len(stack)-1] // pop
				stack = append(stack, num*pre)
			case '/':
				fmt.Println("除")
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num/pre)
			}
			// 😅
			sign, _ = utf8.DecodeRune([]byte{c})
			num = 0
		}
		// 遇到右括号
		if c == ')' {
			break
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
	return char >= 48 && char <= 57
}