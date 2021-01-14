
	/*
		给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

		要求返回这个链表的 深拷贝。

		我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

		val：一个表示 Node.val 的整数。
		random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。

	*/

	// 方法1：回溯。超时
	var s string
	var wordDict []string

	func wordBreak(_s string, _wordDict []string) bool {
		s = _s
		wordDict = _wordDict
		return backtract("")
	}

	func backtract(temp string) bool {
		if len(temp) == len(s) {
			return s == temp
		}
		if len(temp) > len(s) {
			return false
		}
		//
		for i := 0; i < len(wordDict); i++ {
			if backtract(temp + wordDict[i]) {
				return true
			}
		}
		return false
	}

	// 方法2：回溯,优化递归出口。超时
	var s string
	var wordDict []string

	func wordBreak(_s string, _wordDict []string) bool {
		s = _s
		wordDict = _wordDict
		return backtract("")
	}

	func backtract(temp string) bool {
		if len(temp) > len(s) {
			return false
		}
		// 判断此时对应的字符是否全部相等
		for i := 0; i < len(temp); i++ {
			if temp[i] != s[i] {
				return false
			}
		}
		if len(temp) == len(s) {
			return true
		}

		for i := 0; i < len(wordDict); i++ {
			if backtract(temp + wordDict[i]) {
				return true
			}
		}
		return false
	}

	// 方法3：memoization技术。用set判断s和wordDict中的字符是否有不同的，用hash把回溯中已经考虑过的解存起来，第二次回溯过来的时候可以直接使用。
	var s string
	var wordDict []string
	var hash map[string]bool

	func wordBreak(_s string, _wordDict []string) bool {
		s = _s
		wordDict = _wordDict
		hash = make(map[string]bool, 0)
		set := NewSet()
		// 将 wordDict 的每个字母放到 set 中
		for i := 0; i < len(wordDict); i++ {
			word := wordDict[i]
			for j := 0; j < len(word); j++ {
				set.Insert(word[j])
			}
		}
		// 判断 s 的每个字母在 set 中是否存在
		for i := 0; i < len(s); i++ {
			if !set.Contains(s[i]) {
				return false
			}
		}
		return backtract("")
	}
	func backtract(temp string) bool {
		if len(temp) > len(s) {
			return false
		}

		// 之前是否存过
		if val, ok := hash[temp]; ok {
			return val
		}

		for i := 0; i < len(temp); i++ {
			if s[i] != temp[i] {
				return false
			}
		}

		if len(temp) == len(s) {
			return true
		}

		for i := 0; i < len(wordDict); i++ {
			if backtract(temp + wordDict[i]) {
				hash[temp] = true
				return true
			}
		}
		hash[temp] = false
		return false
	}

	// 方法4：分治
	var s string
	var wordDict []string
	var hash map[string]bool
	var set *Set

	func wordBreak(_s string, _wordDict []string) bool {
		s = _s
		wordDict = _wordDict
		hash = make(map[string]bool, 0)
		set = NewSet()
		for i := 0; i < len(wordDict); i++ {
			set.Insert(wordDict[i])
		}
		return fenzhi(s)
	}

	func fenzhi(s string) bool {
		if len(s) == 0 {
			return true
		}
		if val, ok := hash[s]; ok {
			return val
		}

		for i := 0; i < len(s); i++ {
			if set.Contains(s[i:len(s)]) && fenzhi(s[:i]) {
				hash[s] = true
				return true
			}
		}
		hash[s] = false
		return false
	}