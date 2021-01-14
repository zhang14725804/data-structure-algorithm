// 罗马数字转整数
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

func romanToInt2(s string) int {
	// 是有多蠢😅
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'I' {
			if i+1 < len(s) && s[i+1] == 'V' {
				res += 4
				i++
			} else if i+1 < len(s) && s[i+1] == 'X' {
				res += 9
				i++
			} else {
				res += 1
			}
		} else if s[i] == 'X' {
			if i+1 < len(s) && s[i+1] == 'L' {
				res += 40
				i++
			} else if i+1 < len(s) && s[i+1] == 'C' {
				res += 90
				i++
			} else {
				res += 10
			}
		} else if s[i] == 'C' {
			if i+1 < len(s) && s[i+1] == 'D' {
				res += 400
				i++
			} else if i+1 < len(s) && s[i+1] == 'M' {
				res += 900
				i++
			} else {
				res += 100
			}
		} else if s[i] == 'L' {
			res += 50
		} else if s[i] == 'V' {
			res += 5
		} else if s[i] == 'D' {
			res += 500
		} else if s[i] == 'M' {
			res += 1000
		}

	}
	return res
}

func romanToInt(s string) int {
	hash := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	res := 0
	// 注意循环条件，i增长在内部处理
	for i := 0; i < len(s); {
		// 类型转换
		one := string(s[i])
		two := ""
		//
		if i+2 <= len(s) {
			two = string(s[i : i+2])
		}
		if val, ok := hash[two]; ok {
			i = i + 2
			res += val
		} else if val, ok := hash[one]; ok {
			res += val
			i++
		}
	}

	return res
}
