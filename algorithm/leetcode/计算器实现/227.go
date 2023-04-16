/*
	折腾了很久
	😅😅😅
*/
func calculate(s string) int {
	prio := 0 // 优先级
	nums := []int{}
	opers := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '*' || s[i] == '/' {
			prio = 1
			opers = append(opers, s[i])
		} else if s[i] == '+' || s[i] == '-' {
			opers = append(opers, s[i])
		} else {
			j := i
			temp := 0
			for j < len(s) && isdigit(s[j]) {
				temp = temp*10 + int(s[j]-'0')
				j++
			}
			i = j - 1 // 😅
			nums = append(nums, temp)

			// 遇到乘除，nums 在队尾操作
			if prio == 1 {
				prio = 0
				x := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				y := nums[len(nums)-1]
				nums = nums[:len(nums)-1]
				op := opers[len(opers)-1]
				opers = opers[:len(opers)-1]

				nums = append(nums, calc(y, x, op))
			}
		}
	}

	// 遇到加减，nums 在队头顺序操作
	for len(opers) != 0 {
		x := nums[0]
		nums = nums[1:]

		y := nums[0]
		nums = nums[1:]

		op := opers[0]
		opers = opers[1:]

		nums = append([]int{calc(x, y, op)}, nums...)
	}
	return nums[len(nums)-1]
}

func calc(x, y int, op byte) int {
	if op == '-' {
		return x - y
	}
	if op == '+' {
		return x + y
	}
	if op == '/' {
		return x / y
	}
	if op == '*' {
		return x * y
	}
	return -9999
}
func isdigit(char byte) bool {
	return char >= '0' && char <= '9'
}