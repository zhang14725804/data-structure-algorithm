/*
	滑动窗口还是生涩 😅😅😅
*/
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int, 0)
	needSize := 0
	for i := 0; i < len(p); i++ {
		c := p[i]
		need[c]++
		if need[c] == 1 {
			needSize++
		}
	}

	window := make(map[byte]int, 0)
	res := make([]int, 0)
	left, right, valid := 0, 0, 0
	for right < len(s) {
		c := s[right]
		// 移动右指针
		right++
		// 如果c在目标字符串中
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		// 满足窗口大小
		for right-left >= len(p) {
			// 更新答案
			if valid == needSize {
				res = append(res, left)
			}
			// 要删除的字符
			d := s[left]
			// 移动左指针
			left++
			// 如果要删除的字符在目标字符串中
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

/*
	我的滑动窗口，每一次都从零开始 😅😅😅
	超出时间限制
*/
func findAnagrams(s string, p string) []int {
	res := []int{}
	need := make(map[byte]int, 0)
	needTypes := 0
	for _, c := range p {
		bc := byte(c)
		need[bc]++
		if need[bc] == 1 {
			needTypes++
		}
	}

	left := 0
	pLen := len(p)
	right := len(s)
	for left+pLen <= right {
		window := make(map[byte]int, 0)
		curTypes := 0
		for i := left; i < left+pLen; i++ {
			w := s[i]
			window[w]++
			if window[w] == need[w] {
				curTypes++
			}
		}
		if curTypes == needTypes {
			res = append(res, left)
		}
		left++
	}
	return res
}