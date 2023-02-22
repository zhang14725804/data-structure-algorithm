/*
	给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

	思路：双指针，一个从头开始遍历，一个从尾部开始遍历，遇到不相同的情况，分两种情况，如果都不满足返回false
*/

func validPalindrome(s string) bool {
	l,r := 0,len(s)-1
	for l<r{
		// 不同的地方这么处理
		if s[l] != s[r]{
			return isPalindrome(s,l+1,r) || isPalindrome(s,l,r-1)
		}
		l++
		r--
	}
    return true
}

func isPalindrome(s string,l,r int) bool {
	for l<r{
		if s[l] != s[r]{
			return false
		}
		l++
		r--
	}
	return true
}