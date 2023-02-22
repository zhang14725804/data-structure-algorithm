/*

 */
func titleToNumber(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += (int(s[i]) - 'A' + 1) * Pow(26, len(s)-i-1)
	}
	return res
}
