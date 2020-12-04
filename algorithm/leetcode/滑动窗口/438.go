/*
	给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
	字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

	说明：

	字母异位词指字母相同，但排列不同的字符串。
	不考虑答案输出的顺序。
*/
func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)
	left, right, valid, needSize := 0, 0, 0, 0
	for i := 0; i < len(p); i++ {
		c := p[i]
		need[c]++
		if need[c] == 1 {
			needSize++
		}
	}
	res := make([]int, 0)
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		//(question)判断左侧窗口是否要收缩
		for right-left >= len(p) {
			if valid == needSize {
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}