/*
	给定一个单词，你需要判断单词的大写使用是否正确。

	我们定义，在以下情况时，单词的大写用法是正确的：

		全部字母都是大写，比如"USA"。
		单词中所有字母都不是大写，比如"leetcode"。
		如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。

	否则，我们定义这个单词没有正确使用大写字母。

*/
func detectCapitalUse(word string) bool {
	upperLen := 0
	for i := 0; i < len(word); i++ {
		if isUpperCase(word[i]) {
			upperLen++
		}
	}
	// 统计大写字母个数：0，1，len，三种情况分别判断
	if upperLen == 0 || upperLen == len(word) || (upperLen == 1 && isUpperCase(word[0])) {
		return true
	}
	return false
}

func isUpperCase(c byte) bool {
	if c >= 65 && c <= 90 {
		return true
	}
	return false
}
