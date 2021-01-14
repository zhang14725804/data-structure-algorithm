/*
	给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
	说明：本题中，我们将空字符串定义为有效的回文串
*/

// 思路1：筛除字母数字之外的任何字符,然后for循环双指针
func isPalindrome(s string) bool {
	str := ""
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			str += string(ch)
		}

	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if !strings.EqualFold(string(str[i]), string(str[j])) {
			return false
		}
	}
	return true
}

// 思路2：for循环,双指针，跳过字母数字之外的字符
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) && i < j {
			i++
		}
		for !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) && j > i {
			j--
		}
		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
	}
	return true
}