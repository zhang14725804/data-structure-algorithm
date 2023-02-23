/*
	1. 从前向后遍历
		遇到A，count++，判断是否大于2，满足条件return false
		遇到L，for下一个是否是L，count++，判断是否大于等于3，满足条件return false
		否则return true
		我的蠢代码 😅
*/
func checkRecord(s string) bool {
	acount := 0
	for i := 0; i < len(s); i++ {
		// A 没什么好说的
		if s[i] == 'A' {
			acount++
			if acount > 1 {
				return false
			}
		}
		// 判断是否连续三次迟到
		lcount := 0
		for i < len(s) && s[i] == 'L' {
			lcount++
			i++
		}
		if lcount > 2 {
			return false
		} else if lcount > 0 {
			// 注意这里要退回1位 "ALLAPPL"
			i--
		}

	}
	return true
}

func checkRecord(s string) bool {
	acount := 0
	late := 0
	for _, c := range s {
		if c == 'A' {
			acount++
			if acount > 1 {
				return false
			}
		}
		if c == 'L' {
			late++
			if late > 2 {
				return false
			}
		} else {
			// 不是L切没有满足犯错条件
			late = 0
		}
	}
	return true
}