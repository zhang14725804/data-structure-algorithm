/*
	给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
	如果不存在最后一个单词，请返回 0 。
*/
func lengthOfLastWord(s string) int {
	// 直接从最后一个字符往前遍历，遇到空格停止就可以了
	index := len(s) - 1
	// 过滤到末尾的空格
	for {
		if index < 0 || s[index] != ' ' {
			break
		}
		index--
	}
	count := 0
	for i := index; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}