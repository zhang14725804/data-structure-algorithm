/*
	折腾了很久
	😅😅😅
*/

func calculate(s string) int {
	sign := 1 // 默认【+】
	sum := 0
	stk := []int{} // 存储当前sum和sign

	sl := len(s)

	for i := 0; i < sl; {
		ch := s[i]
		if ch == ' ' { // 空格
			i++
			continue
		} else if ch == '-' { // 加减号，改变sign
			i++
			sign = -1
		} else if ch == '+' {
			i++
			sign = 1
		} else if ch == '(' { // 左括号
			// 将当前sum和sign入栈
			stk = append(stk, sum)
			stk = append(stk, sign)
			// 初始化sign和sum
			sign = 1
			sum = 0
			i++
		} else if ch == ')' { // 右括号
			// 将上一个sum和sign出栈
			preSign := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			preSum := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			// 计算结果
			sum = preSum + preSign*sum
			i++
		} else { // 数字
			j := i
			temp := 0
			for j < len(s) && isdigit(s[j]) {
				temp = temp*10 + int(s[j]-'0')
				j++
			}
			// 更新当前sum
			sum += temp * sign
			i = j
		}
	}
	return sum
}