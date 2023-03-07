/*
	方法1:
		1. 遍历字符串数组，将字符串排序作为key存入hash，当前字符串作为value
		2. 最后遍历hash
*/
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, 0)
	for _, str := range strs {
		b := []byte(str)
		// 字符串内部排序，排序后相同的结果存入map
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		hash[string(b)] = append(hash[string(b)], str)
	}
	res := make([][]string, 0)
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}

/*
	方法2:
		1. 声明一个26位长的数组asciiArray，初始值为0
		2. 遍历字符串数组，每个字符串的字符取ASCII然后减去97，对应asciiArray位置上++；
		3. 数组转换为字符串作为key，存入hash，当前字符串存入value（这个value是数组）
		4. 遍历hash，返回结果；
*/
func groupAnagrams(strs []string) [][]string {
	strCnt := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		ascii := make([]int, 26)
		for j := 0; j < len(strs[i]); j++ {
			idx := strs[i][j] - 97
			ascii[idx]++
		}
		// array转换为字符串
		strAsc, _ := json.Marshal(ascii)
		strCnt[string(strAsc)] = append(strCnt[string(strAsc)], strs[i])
	}

	res := make([][]string, 0)
	for _, v := range strCnt {
		res = append(res, v)
	}
	return res
}