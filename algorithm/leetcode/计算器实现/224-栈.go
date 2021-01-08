//

func calculater(s string) int {
	oper := make([]byte, 0)
	nums := make([]int, 0)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		// 空格
		if ch == ' ' {
			continue
		}
		// 符号和左括号
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			oper = append(oper, ch)
		} else if ch == '(' {
			oper = append(oper, ch)
		} else if ch == ')' {
			// 右括号
			if len(oper) > 0 && oper[len(oper)-1] != '(' {
				calc(&oper, &nums)
			}
			oper = oper[:len(oper)-1] // 这一步什么意思
		} else {
			j := i
			// 字符串转数字
			x := 0
			for j < len(s) && isdigit(s[j]) {
				x = x*10 + int(s[j]-'0')
				j++
			}
			nums = append(nums, x)
			// 回退一步
			i = j - 1
			if len(oper) > 0 && oper[len(oper)-1] != '(' {
				calc(&oper, &nums)
			}
		}
	}
	// fmt.Println(oper, nums)
	if len(oper) > 0 {
		calc(&oper, &nums)
	}
	return nums[len(nums)-1]
}

// question 指针 & 地址问题
func calc(oper *[]byte, nums *[]int) {
	// type *[]int does not support indexing
	// top & pop
	y := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	// fmt.Println("=====", nums, y)
	x := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	// fmt.Println("*****", nums, x)
	sign := (*oper)[len(*oper)-1]
	*oper = (*oper)[:len(*oper)-1]
	if sign == '+' {
		*nums = append(*nums, x+y)
	}
	if sign == '-' {
		*nums = append(*nums, x-y)
	}
	if sign == '*' {
		*nums = append(*nums, x*y)
	}
	if sign == '/' {
		*nums = append(*nums, x/y)
	}
}
func isdigit(char byte) bool {
	return char >= '0' && char <= '9'
}

// func main() {
// 	fmt.Println(calculater("1+(4+5+2) - 3"))
// 	fmt.Println(calculater("(1+(4+5)-3)+(6+8)"))
// 	fmt.Println(calculater(" ( 5*(4-2)+1-6)/5*2 "))
// 	fmt.Println(calculater(" (12-1) + 12 "))
// 	fmt.Println(calculater("(12/2+3)"))
// 	fmt.Println(calculater("-2+1"))  // todo question 有问题😅
// }
