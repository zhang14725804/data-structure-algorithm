package main

/*
	给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

	思路：滑动窗口(todo：思路和代码有点绕)
*/
func lengthOfLongestSubstring(s string) int {
	n, ans := len(s), 0
	hash := make(map[byte]int) // 存放字符出现的位置
	for i, j := 0, 0; j < n; j++ {
		if _, ok := hash[s[j]]; ok {
			// 发现重复的，则重新选一个i，这个i停留再出现重复的下一位置
			i = compare(hash[s[j]], i, true)
		}
		// 更新最大值
		ans = compare(ans, j-i+1, true)
		hash[s[j]] = j + 1
	}
	return ans
}
func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}

	return a
}
/*
	双指针算法
*/
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	res :=0
	for i,j := 0,0; i<len(s); i++{
		// s[i]存入hash中
		hash[s[i]]++
		// hash[s[i]] > 1的时候，hash[s[i]] == hash[s[j]]
		for hash[s[i]] > 1 {
			// hash[s[i]]-- 同样可以
			hash[s[j]]--
			// j向前移动一位
			j++
		}
		res = compare(res, i-j+1, true)
	}
	return res
}
