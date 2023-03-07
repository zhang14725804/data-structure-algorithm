/*
	么得方法，硬上😅
	1. 初始字符串s=“1”
	2. 循环n-1次
	3. 遍历当前字符串
	4. 统计相同的元素的个数，更新字符串
*/
func countAndSay(n int) string {
	s := "1"
	// 循环n-1次
	for i := 0; i < n-1; i++ {
		ns := ""
		// 遍历当前字符串，注意，后边不是j++
		for j := 0; j < len(s); {
			// k=j+1而不是k=0或者k=j
			k := j + 1
			// 统计相同的元素
			for k < len(s) && s[k] == s[j] {
				k++
			}
			// +=而不是=
			ns += fmt.Sprintf("%d%s", k-j, string(s[j]))
			// 更新k的位置
			j = k
		}
		s = ns
	}
	return s
}