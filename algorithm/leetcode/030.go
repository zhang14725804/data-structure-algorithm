/*
	给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

	注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

	todo：滑动窗口？不会做
*/

func findSubstring(s string, words []string) []int {
	windowLen := len(words[0])
	hash := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		hash[words[i]] = 0
	}
	sign := 0
}