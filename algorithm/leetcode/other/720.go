/*
	没懂题意 😅
	1. 将words数组排序。首先按照单词的长度升序排序，如果单词的长度相同则按照字典序降序排序
	2. 排序之后，可以确保当遍历到每个单词时，比该单词短的全部单词都已经遍历过，且每次遇到符合要求的单词一定是最长且字典序最小的单词，可以直接更新答案。
	3. 将答案初始化为空字符串。使用哈希集合存储所有符合要求的单词，初始时将空字符串加入哈希集合
	4. 对于每个单词，判断当前单词去掉最后一个字母之后的前缀是否在哈希集合中，如果该前缀在哈希集合中则当前单词是符合要求的单词，将当前单词加入哈希集合，并将答案更新为当前单词
*/
func longestWord(words []string) (longest string) {
	// 排序
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	//
	cache := map[string]struct{}{"": {}}
	for _, word := range words {
		//
		if _, ok := cache[word[:len(word)-1]]; ok {
			longest = word
			cache[word] = struct{}{}
		}
	}
	return longest
}
