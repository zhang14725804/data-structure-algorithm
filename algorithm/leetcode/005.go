package leetcode

import "fmt"

/*
	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	示例 1：
	输入: "babad"
	输出: "bab"
	注意: "aba" 也是一个有效答案。

	示例 2：
	输入: "cbbd"
	输出: "bb"

	ps：中心扩展法和动态规划法
	题目的标签是动态规划，关键点是，
	当s[i,j]是回文串 且 s[i-1] == s[j+1]的时候 s[i-1, j+1] 也是回文串， 定义一个dp数组， 初始化dp[i][i] 与 dp[i][i+1]，然后动态规划
*/
// Leetcode005 动态规划(有问题)
func Letcode005(s string) string {
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
	fmt.Println(dp)
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
	fmt.Println(s[left:(right - left + 1)])
	return s[left:(right - left + 1)]
}
