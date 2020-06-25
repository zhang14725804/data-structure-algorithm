/*
	3. Longest Substring Without Repeating Characters

	暴力是怎么写的，然后在对暴力做法进行优化
	todo：思路理解：双指针法（同时向一个方向移动）

	todos::"abcabcbb"测试通过，"pwwkew"无法通过
*/
func lengthOfLongestSubstring(str string) int {
	s:=[]rune(str)
	hash:=make(map[rune]int)
	res:=0
	// i指针在前，j指针在后
	for i,j:=0,0; i<len(s); i++{
		// 如果有重复的元素，一定是s[i]
		hash[s[i]]+=1
		// 有重复元素
		for hash[s[i]] > 1{
			j+=1
			hash[s[j]]-=1
			res = max(res, i-j+1)
		}
	}
	return res
}


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}