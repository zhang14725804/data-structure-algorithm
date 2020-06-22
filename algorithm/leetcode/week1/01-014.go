/*
	编写一个函数来查找字符串数组中的最长公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}

	res := ""
	for i:=0; ;i++{
		flag := false
		first := strs[0]
		for _,now := range strs{
			// 临界条件
			if i>= len(first) || i >= len(now) || first[i] != now[i]{
				flag = true
				break
			}
		}
		if flag == true{
			break
		}
		// 注意：first[i]是byte类型
		res += string(first[i])
	}
	return res
}