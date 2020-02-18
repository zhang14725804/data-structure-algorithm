package leetcode
/*
	49. Group Anagrams

	如何发现乱序字符串的本质
*/
func groupAnagrams(strs []string) [][]string {
	// 声明hash
	hash:= make(map[string][]string,0)
	// 重点在这里
	for _, str := range strs { 
		b := []byte(str)
		// 利用将两个字符串排序，来判断两个字符串是不是异位词
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		hash[string(b)] = append(hash[string(b)], str) 
	}
	// 声明数组
	res := make([][]string,0)
	for _, v := range hash {
		res = append(res,v)
	}
	return res
}