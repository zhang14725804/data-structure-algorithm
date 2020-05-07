/*
	最长不含重复字符的子字符串
	思路：双指针问题
	todo："chbmmcg"测试不通过
*/
func longestSubstringWithoutDuplication(s string) int {
	count := make(map[byte]int)
	res:=0
	for i,j:=0,0; j<len(s); j++{
		count[s[j]]++
		if count[s[j]] > 1 {
			for count[s[i]] > 1{
				count[s[i]]--
				i++
			}
			i++
			count[s[i]]--
		}
		res = max(res, j-i+1)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}