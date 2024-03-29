/*
	方法1：滑动窗口
		1. 统计t中字符的种类和每个字符出现的次数
		2. 双指针left,right
		3. 移动right，检查当前元素是否在t中
		4. 若存在，更新窗口内容，检查是否满足t元素种类和个数
		5. 若满足要求，右移left收缩窗口，收缩的过程中更新窗口内容
*/
func minWindow(s string, t string) string {
	need := make(map[byte]int, 0)
	// needSize 要满足的元素个数
	needSize := 0
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
		if need[c] == 1 {
			// 要满足的元素种类
			needSize++
		}
	}

	// valid 当前已满足的元素个数
	left, right, valid := 0, 0, 0, 0
	window := make(map[byte]int, 0)
	var res string
	for right < len(s) {
		// 将要入窗口的字符
		c := s[right]
		// 移动右边指针
		right++
		// 若当前字符存在于t中
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 😅 持续收缩，判断左侧边界是否需要收缩
		for valid == needSize {
			// 更新最小覆盖字串
			if len(res) == 0 || len(res) > right-left {
				res = s[left:right]
			}
			// 将要移除窗口的字符
			d := s[left]
			// 向左移动左边指针 😅😅😅 为什么不需要判断left是否越界
			left++
			// 若当前字符存在于t中
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

/*
	方法2：滑动窗口（优化版）(question) 😅😅😅
		TODO
*/
func minWindow2(s1 string, t1 string) string {
	// need始终记录着当前滑动窗口下，我们还需要的元素种类
	need := make(map[byte]int, 0)
	// 记录所需元素种类
	cnt := 0
	// 字符串遍历：range遍历的结果是rune，for便利的结果是byte
	for i := 0; i < len(t1); i++ {
		c := t1[i]
		need[c]++
		if need[c] == 1 {
			cnt++
		}
	}

	var res string

	// 数量为负的就是不必要的元素，而数量为1表示刚刚好
	// 从前往后维护窗口。i先走，j在后
	for i, j, count := 0, 0, 0; i < len(s1); i++ {
		// (question)没看懂；为什么count只加不减，need中有加有减，count为什么就特殊了😅
		// 满足当前字母和个数
		if need[s1[i]] == 1 {
			count++
		}
		// 需要的字母个数减减
		need[s1[i]]--
		// 当前字母无关紧要(数量为负的就是不必要的元素)，j可以向后移动
		for j < i && need[s1[j]] < 0 {
			need[s1[j]]++
			j++
		}
		// 如何判断滑动窗口包含了T的所有元素？结论就是当need中所有元素的数量都小于等于0时，表示当前滑动窗口不再需要任何元素
		if count == cnt {
			if len(res) == 0 || len(res) > i-j+1 {
				res = s1[j : i+1]
			}
		}
	}
	return res
}