/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

*/

/*
	滑动窗口，优化版(question)（看不懂😅）
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

/*
	滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int, 0)
	left, right, res := 0, 0, 0
	for right < len(s) {
		c := s[right]
		// 移动右指针
		right++
		window[c]++
		// 重复的情况
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		// 在收缩窗口完成后更新 res，因为窗口收缩的 while 条件是存在重复元素，换句话说收缩完成后一定保证窗口中没有重复嘛
		res = MaxInt(res, right-left)
	}
	return res
}