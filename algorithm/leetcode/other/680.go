/*
	思路：双指针
	start从头开始遍历，end从尾部开始遍历，遇到不相同的情况：
	1. 跳过start，和end继续
	2. 跳过end，和start继续
	如果都不满足返回false
*/

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		// 😅 不同的地方这么处理
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}