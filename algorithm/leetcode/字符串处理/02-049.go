/*
	给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

	如何发现乱序字符串的本质
*/
func groupAnagrams(strs []string) [][]string {
	// 声明hash
	hash := make(map[string][]string, 0)
	// 重点在这里
	for _, str := range strs {
		b := []byte(str)
		// 字符串内部排序，排序后相同的结果存入map
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		hash[string(b)] = append(hash[string(b)], str)
	}
	//
	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}