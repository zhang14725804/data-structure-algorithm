/*
	3. Longest Substring Without Repeating Characters

	双指针法（同时向一个方向移动）,思路精巧
	todos::"abcabcbb"测试通过，"pwwkew"无法通过
*/
func lengthOfLongestSubstring(str string) int {
	s:=[]rune(str)
	hash:=make(map[rune]int)
	res:=0
	for i,j:=0,0;i<len(s);i++{
		hash[s[i]]+=1
		for hash[s[i]] > 1{
			j+=1
			hash[s[j]]-=1
			res = max(res,i-j+1)
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