
/*
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

	说明：
	分隔时可以重复使用字典中的单词。
	你可以假设字典中没有重复的单词。
*/

// 思路：分治
var s string
var set *Set
var hash map[string][]string
var wordDict []string

func wordBreak(_s string, _wordDict []string) []string {
	s = _s
	wordDict = _wordDict
	hash = make(map[string][]string, 0)
	set = NewSet()

	for i := 0; i < len(wordDict); i++ {
		set.Insert(wordDict[i])
	}

	return fenzhi(s)
}

// todo：没懂😅
func fenzhi(s string) []string {
	var res []string
	if len(s) == 0 {
		return res
	}

	if val, ok := hash[s]; ok {
		return val
	}

	for i := 0; i < len(s); i++ {
		// 判断当前字符串是否存在
		if set.Contains(s[i:len(s)]) {
			//
			if i == 0 {
				res = append(res, s[i:len(s)])
			} else {
				temp := fenzhi(s[:i])
				for j := 0; j < len(temp); j++ {
					res = append(res, temp[j]+" "+s[i:len(s)])
				}
			}
		}
	}
	hash[s] = res
	return res
}
