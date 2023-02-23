/*
	1. 从前向后走，记录位置i
	2. 遍历数组
		a. 记录第一个字符串的第一个字符
		b. 如果i长度大于第一个字符串的长度
			或者i大于当前遍历字符串的长度
			或者第一个字符串的i和当前字符串的i未知元素不同退出

	么的思路 😅😅😅
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := ""
	// 注意循环条件 😅
	for i := 0; ; i++ {
		flag := false
		first := strs[0]
		for _, now := range strs {
			// 注意临界条件 😅
			if i >= len(first) || i >= len(now) || first[i] != now[i] {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		res += string(first[i])
	}
	return res
}