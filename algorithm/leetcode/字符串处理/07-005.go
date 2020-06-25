/*
	5. Longest Palindromic Substring(最长回文子串)

	马拉车算法

	todos：："cbbd"、"bb"测试不通过！！！
*/
func longestPalindrome(s string) string {
    var res string
	for i:=0;i<len(s);i++{
		// 一个对称点（字符串长度是奇数），由中心点往两边走（j往左，k往右）
		for j,k:=i,i; j>=0 && k <len(s) && s[j] == s[k]; j,k=j-1,k+1{
			if len(res) < k-j+1{
				res = s[j:k-j+1]
			}
		}
		// 两个对称点（字符串长度是偶数），由中心点往两边走（j往左，k往右）
		for j,k:=i,i+1; j>=0 && k <len(s) && s[j] == s[k]; j,k = j-1,k+1{
			if len(res) < k-j+1{
				res = s[j:k-j+1]
			}
		}
	}
	return res
}