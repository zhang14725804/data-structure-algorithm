func countAndSay(n int) string {
	s := "1"
	// 循环n次
	for i := 0; i < n-1; i++ {
		ns := ""
		// 遍历当前字符串
		for j := 0; j < len(s); j++ {
			// 统计相同的元素
			k := j
			for k<len(s) && s[k] == s[j]{
				k++
			}
			// 注意：不能直接用string(k-j)
			ns += strconv.Itoa(k-j) + string(s[j])
			// 更新k的位置
			j = k-1
		}
		s = ns
	}
	return s
}