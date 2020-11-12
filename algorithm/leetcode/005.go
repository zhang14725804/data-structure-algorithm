package main

/*
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	思路1:双指针
	思路2：回文字符串专用算法--马拉车*/

// todo:"cbbd","bb","babad",无法同时通过测试
func longestPalindrome(s string) string {
	res := ""
	// 遍历每个字符，以每个字符为中心，像两边扩散
	for i := 0; i < len(s); i++ {
		// 数组长度是奇数
		l, r := i-1, i+1
		// 两个指针同时向左向右背向移动
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l-1 {
			res = s[l+1 : r-l-1]
		}
		// 数组长度是偶数
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l-1 {
			res = s[l+1 : r-l]
		}
	}
	return res
}
