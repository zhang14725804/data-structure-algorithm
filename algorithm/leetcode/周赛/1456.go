/*
	给你字符串 s 和整数 k 。

	请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。

	英文中的 元音字母 为（a, e, i, o, u）。

	思路：双指针（滑动窗口）。
*/
func maxVowels(s string, k int) int {
	res := 0
	cnt := 0
	// //移动右端点 i
	for i := 0; i < len(s); i++ {
		cnt += check(s[i])
		//窗口超过 k 了，
		if i >= k {
			//从窗口中移除s[j], j = i-k
			cnt -= check(s[i-k])
		}
		// 更新最大值
		res = compare(res, cnt, true)
	}
	return res
}

//
func check(c byte) int {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return 1
	}
	return 0
}
