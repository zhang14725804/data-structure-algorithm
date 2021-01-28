/*
	给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的那一个元素
	所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

	说明:

	s 可能为空，且只包含从 a-z 的小写字母。
	p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
*/

/*
	暴力递归；可以增加memory缓存优化递归
*/
func isMatch1(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	// 第一个字符是否匹配（字符相等或者'.'）
	firstMatch := len(s) > 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		// 匹配0次（p 跳过两个字符，*前面的字符出现0次） || (第一个字符匹配 && 跳过该字符继续匹配剩余部分)
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}
	return firstMatch && isMatch(s[1:], p[1:])
}