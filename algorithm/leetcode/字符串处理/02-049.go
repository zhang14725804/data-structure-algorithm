/*
	给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

	思路1：遍历字符串数组，将字符串排序作为key存入hash，当前字符串作为value，最后遍历hash
	思路2：声明一个26位长的数组asciiArray，初始值为0，遍历字符串数组，每个字符串的字符取ASCII然后减去97，对应asciiArray位置上++；
	数组转换为字符串作为key，存入hash，当前字符串存入value（这个value是数组），遍历hash，返回结果；（其实和思路1一样）
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