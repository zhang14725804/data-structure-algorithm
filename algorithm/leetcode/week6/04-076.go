/*
	inimum Window Substring
	给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

	（1）暴力方法
	（2）滑动窗口（不好理解），特殊的双指针问题
	如何寻找单调性
	todo：代码有问题。不好理解
*/
func minWindow(s string, t string) string {
	s1 := []rune(s)
	t1 := []rune(t)
	hash := make(map[rune]int, 0)
	// 字符串遍历：range遍历的结果是rune，for便利的结果是byte
	for _, c := range t1 {
		hash[c]++
	}
	// 统计hash表长度（办法好像有点蠢）
	cnt := 0
	for _, _ = range hash {
		cnt++
	}

	var res []rune
	// 从前往后维护窗口。后一个指针i，走在前面的
	for i, j, c := 0, 0, 0; i < len(s1); i++ {
		// 当前字母个数是1
		if hash[s1[i]] == 1 {
			c++
		}
		// 需要的字母个数减减
		hash[s1[i]]--
		// 当前字母无关紧要，j可以向后移动
		// 前一个指针j，走在后面的
		for hash[s1[j]] < 0 {
			// hash[s1[j++]]++
			hash[s1[j]]++
			j++
		}
		// 所有字母都满足要求，更新答案
		if c == cnt {
			if len(res) == 0 || len(res) > i-j+1 {
				res = s1[j : i-j+1]
			}
		}
	}
	return string(res)
}