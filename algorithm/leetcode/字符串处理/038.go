package leetcode
/*
	38. Count and Say

	题目比较难理解
*/
func countAndSay(n int) string {
	s := "1"
	for i:=0;i<n-1;i++{
		ns:=""
		for j:=0;j<len(s);j++{
			k:=j
			for k<len(s) && s[k]==s[j] {
				k++
			} 
			/*
				invalid operation: strconv.Itoa(k - j) + s[j] (mismatched types string and byte) (solution.go)
				+= strconv.Itoa(k-j)+s[j]
				s[j]是byte类型
			*/
			ns += strconv.Itoa(k-j) + string(s[j])
			j=k-1
		}
		s = ns
	}
	return s
}