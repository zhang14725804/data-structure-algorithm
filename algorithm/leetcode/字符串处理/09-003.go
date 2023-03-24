/*
	我的滑动窗口 😅
*/
func lengthOfLongestSubstring(s string) int {
	window := map[byte]int{}
	left := 0
	right := 0
	res := 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		// 没有重复字符 可以把这一步挪到最后，就不用判断了，蠢😅
		if window[c] == 1 {
			res = Max(res, right-left)
			continue
		}

		// 有重复字符
		if window[c] > 1 {
			// 从左向右删
			for window[c] > 1 {
				window[s[left]]--
				left++
			}
		}
	}
	return res
}

/*
	稍微有些进步思想的滑动窗口
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
	方法1:滑动窗口
	1. 指针start,end，一个在前，一个在后，hash记录出现的字符
	2. 从前向后遍历，若当前字符不存在，更新res，end向前走
	3. 若当前字符存在，删除start字符，start向前走
*/
func lengthOfLongestSubstring(s string) int {
	set := map[byte]struct{}{}
	res := 0
	start, end := 0, 0
	for start < len(s) && end < len(s) {
		if _, ok := set[s[end]]; !ok {
			set[s[end]] = struct{}{}
			res = Max(res, end-start+1)
			end++
		} else {
			// 😅 删除start字符，start向前走
			delete(set, s[start])
			start++
		}
	}
	return res
}

/*
	方法2: 滑动窗口优化 😅😅😅
	set 改为 map ，将字符存为 key ，将对应的下标存到 value
*/
func lengthOfLongestSubstring(s string) int {
	hash := map[byte]int{}
	res := 0
	start, end := 0, 0
	for ; end < len(s); end++ {
		// 有重复字符，更新start位置 😅
		if idx, ok := hash[s[end]]; ok {
			start = Max(idx, start)
		}
		res = Max(res, end-start+1)
		// 😅😅😅 下标 + 1 代表 start 要移动的下个位置
		// 为什么要【end + 1】
		hash[s[end]] = end + 1
	}
	return res
}