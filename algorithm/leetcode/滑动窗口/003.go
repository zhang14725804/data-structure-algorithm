/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

/*
	滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	left, right, res := 0, 0, 0
	for right < len(s) {
		cw := s[right] // current word
		// 移动右指针
		right++
		window[cw]++
		// 存在重复元素，收缩窗口【左边】，直到没有重复的元素位置
		for window[cw] > 1 {
			dw := s[left] // delete word
			left++
			window[dw]--
		}
		// 在收缩窗口完成后更新
		res = MaxInt(res, right-left)
	}
	return res
}

/*
	滑动窗口，优化版
	question （看不懂😅）
*/
func lengthOfLongestSubstring1(s string) int {
	n, ans := len(s), 0
	hash := make(map[byte]int) // 存放字符出现的位置
	// i，j分别代表左右指针
	for i, j := 0, 0; j < n; j++ {
		c := s[j]
		// 发现重复的，则重新选一个i，这个i停留再出现重复的下一位置
		if _, ok := hash[c]; ok {
			i = MaxInt(hash[c], i)
		}
		// 更新最大值
		ans = MaxInt(ans, j-i+1)
		// ？？？
		hash[c] = j + 1
	}
	return ans
}