/*
	1. 统计所有字符的个数
	2. 统计结果，偶数直接相加，奇数减一相加
	3. 如果存在奇数个字符，最终结果加一
*/
func longestPalindrome(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	sb := []byte(s)
	cmap := make(map[byte]int, 0)
	for _, c := range sb {
		cmap[c]++
	}

	res := 0
	exitOdd := false
	for _, ct := range cmap {
		if ct%2 == 0 {
			res += ct
		} else {
			res += ct - 1
			exitOdd = true
		}
	}
	if exitOdd {
		res++
	}
	return res
}