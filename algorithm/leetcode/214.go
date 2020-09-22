/*
	最短回文串

	方法1：暴力破解。先判断整个字符串是不是回文串，如果是的话，就直接将当前字符串返回。不是的话，进行下一步。
	判断去掉末尾 1 个字符的字符串是不是回文串，如果是的话，就将末尾的 1 个字符加到原字符串的头部返回。不是的话，进行下一步

	方法2：将原始字符串逆序，然后比较对应的子串即可判断是否是回文串，将末尾不是回文串的部分倒置加到原字符串开头即可。

	todo：六种解法😅，这么优秀嘛
*/

// 采用方法2,骚气
func shortestPalindrome(s string) string {
	i, n := 0, len(s)
	r := string(reverse([]rune(s)))

	for ; i < n; i++ {
		if s[:n-i] == r[i:] {
			break
		}
	}
	return string(reverse([]rune(s[n-i:]))) + s
}