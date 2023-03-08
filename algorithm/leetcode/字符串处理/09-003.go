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
		// 更新start位置 😅
		if idx, ok := hash[s[end]]; ok {
			start = Max(idx, start)
		}
		res = Max(res, end-start+1)
		// 😅😅😅 下标 + 1 代表 start 要移动的下个位置
		hash[s[end]] = end + 1
	}
	return res
}