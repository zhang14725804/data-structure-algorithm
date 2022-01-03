/*
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
	思路1:双指针
	思路2：回文字符串专用算法--马拉车
	思路3：动态规划
*/

/*
	思路1： 双指针
		（1）以【每个】字母为中心向两边扩散
		（2）以【每两个】字母为中心向两边扩散
	0102 还是做错（也就是不会）
*/
func longestPalindrome1(s string) string {
	res := ""
	// 遍历每个字符，以每个字符为中心，像两边扩散
	for i := 0; i < len(s); i++ {
		// (1) 找到以 s[i] 为中心的回文串
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l {
			res = s[l+1 : r]
		}
		// (2) 找到以 s[i] 和 s[i+1] 为中心的回文串
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l {
			res = s[l+1 : r]
		}
	}
	return res
}

/*
	动态规划
	当s[i,j]是回文串 且 s[i-1] == s[j+1]的时候 s[i-1, j+1] 也是回文串， 定义一个dp数组， 初始化dp[i][i] 与 dp[i][i+1]
*/
func longestPalindrome(s string) string {
	sLen := len(s)
	// 初始化判断
	if sLen < 2 {
		return s
	}
	// 初始化动态规划dp数组，dp[i][j]表示从j到i的字符串是否为回文串
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}
	// 初始化最长的最优节点
	right, left := 0, 0
	for i := 0; i < sLen; i++ {
		// 只有一个元素的时候肯定为true
		dp[i][i] = true
		// 遍历到第i个元素，再倒退判断是否为回文串
		for j := i - 1; j >= 0; j-- {
			// 头i尾j两个元素相等，且dp[i-1][j+1]是回文串，即dp[i][j]也是回文串
			// 特殊情况,“bb”,此时dp[i-1][j+1]=dp[j][i]此时数组不成立，不存在截取的反向字符串
			dp[i][j] = (s[i] == s[j]) && ((i-j) == 1 || dp[i-1][j+1])
			//截取的字符串的长度大于之前求得的左右长度，则取的左右下标点
			if dp[i][j] == true && (i-j) > (right-left) {
				right = i
				left = j
			}
		}
	}
	return s[left:(right + 1)]
}
