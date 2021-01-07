package main

import (
	"fmt"
	"unicode/utf8"
)

// func main() {
// fmt.Println(calculate("1+(4+5+2) - 3"))
// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
// fmt.Println(calculate(" ( 5*(4-2)+1-6)/5*2 "))
// fmt.Println(calculate(" (12-1) + 12 "))
// 	fmt.Println(calculate("(12-1)"))
// }

/*
	实现一个基本的计算器来计算一个简单的字符串表达式的值。
	字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，乘号 *，除号 /，非负整数和空格  。

*/

/*
	question 【处理复杂问题的思路】
	（1）字符串转数字
	（2）处理只包含加减法的算式
	（3）处理包含加减乘除四则运算的算式
	（4）处理空格字符
	（5）处理包含括号的算式
*/
//
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	return helper(s)
}

// question 遇到括号递归有问题
func helper(s string) int {
	stack := make([]int, 0)
	num := 0
	sign := '+'
	for len(s) > 0 {
		c := s[0]
		s = s[1:]
		if isdigit(c) {
			num = num*10 + int(c-'0')
		}
		if c == '(' {
			num = helper(s)
		}
		if (!isdigit(c) && c != ' ') || len(s) == 0 {
			if sign == '-' {
				stack = append(stack, -num)
			} else if sign == '+' {
				stack = append(stack, num)
			} else if sign == '*' {
				pre := stack[len(stack)-1]   // top
				stack = stack[:len(stack)-1] // pop
				stack = append(stack, num*pre)
			} else if sign == '/' {
				pre := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, num/pre)
			}

			sign, _ = utf8.DecodeRune([]byte{c})
			num = 0
		}
		if c == ')' {
			break
		}
	}
	fmt.Println(stack)
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
	// char >= 48 && char <= 57
	return char >= '0' && char <= '9'
}
