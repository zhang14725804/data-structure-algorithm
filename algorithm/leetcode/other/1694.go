

func reformatNumber(number string) string {
	// 去除非数字
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	ans := []string{}
	i := 0
	// 大于4的直接加入
	for ; i+4 < len(s); i += 3 {
		ans = append(ans, s[i:i+3])
	}
	s = s[i:]
	// 小于4的部分；等于4的部分
	if len(s) < 4 {
		ans = append(ans, s)
	} else {
		ans = append(ans, s[:2], s[2:])
	}
	// 组合答案
	return strings.Join(ans, "-")
}

/*
	1. 从前向后遍历，去掉非字符
	2. 标记是否满足5个号码，标记上一个破折号的位置
	3. 满5个号码，标记加一个破折号，更新破折号标记
	我的是这个样子  😅😅😅
*/
func reformatNumber(number string) string {
	number := "123 4-5678"
	faddr := 0
	fnum := 0
	res := make([]byte, 0)

	for i := 0; i < len(number); i++ {
		if number[i] > 47 && number[i] < 58 {
			res = append(res, number[i])
			fnum++
		}
		if fnum > 4 {
			faddr += 3
			// 😅😅😅 这里拼接数组还折腾了半天
			// res = append(append(res[:faddr], '-'), res[faddr:]...)
			res = append(res[:faddr], append([]byte{'-'}, res[faddr:]...)...)
			fnum -= 4
		}
	}
	if fnum == 4 {
		faddr += 3
		res = append(res[:faddr], append([]byte{'-'}, res[faddr:]...)...)
	}
	return string(res)
}