/*
	给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有 字符 的最小子串。
	滑动窗口（不好理解），特殊的双指针问题
*/

/*
	方法1：暴力做法
	枚举所有字串，检查是否包含所有T，找到最小字串
*/

/*
	方法2：滑动窗口
	如何快速判断当前字串是否包含t中所有字母，hash或数组
*/
func minWindow(s1 string, t1 string) string {
	// need始终记录着当前滑动窗口下，我们还需要的元素数量
	// 数量为负的就是不必要的元素，而数量为0表示刚刚好
	need := make(map[byte]int, 0)
	// 记录所需元素的总数量（不同字母的种类）
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