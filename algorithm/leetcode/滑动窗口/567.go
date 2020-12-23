/*
	给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
	换句话说，第一个字符串的排列之一是第二个字符串的子串。

	输入的字符串只包含小写字母
	两个字符串的长度都在 [1, 10,000] 之间
*/

/*
	滑动窗口
	(question)
*/
func checkInclusion(t string, s string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	needSize, left, right, valid := 0, 0, 0, 0
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
		if need[c] == 1 {
			needSize++
		}
	}

	for right < len(s) {
		c := s[right]
		right++
		// 对窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩。窗口大小大于len(t)时，可以移动left缩小窗口
		// (question)移动 left 缩小窗口的时机是窗口大小大于 t.size() 时，应为排列嘛，显然长度应该是一样的。
		for right-left >= len(t) {
			// 当发现 valid == need.size() 时，就说明窗口中就是一个合法的排列，所以立即返回 true
			if valid == needSize {
				return true
			}
			d := s[left]
			left++
			// 对窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}