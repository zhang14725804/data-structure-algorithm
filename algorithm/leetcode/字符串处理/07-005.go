/*
	马拉车算法
		1. 选一个点😅向两边扩散【j, k := i, i】
		2. 选两个点😅😅向两边扩散 【j, k := i, i+1】
		3. 若符合条件，更新结果s[j : k+1]
*/
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 【j, k := i, i】 😅一个对称点（字符串长度是奇数），由中心点往两边走（j往左，k往右）
		for j, k := i, i; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
			if len(res) < k-j+1 {
				res = s[j : k+1]
			}
		}
		// 【j, k := i, i+1】😅两个对称点（字符串长度是偶数），由中心点往两边走（j往左，k往右）
		for j, k := i, i+1; j >= 0 && k < len(s) && s[j] == s[k]; j, k = j-1, k+1 {
			if len(res) < k-j+1 {
				res = s[j : k+1]
			}
		}
	}
	return res
}