/*
	给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
	返回符合要求的最少分割次数。
*/

/*
	方法1：【分治】+memoization
	（可以改为动态规划 😅😅😅😅😅😅 ）
	大问题化小问题，利用小问题的结果，解决当前大问题。

	aabb
	先考虑在第 1 个位置切割，a | abb
	这样我们只需要知道 abb 的最小切割次数，然后加 1，记为 m1

	aabb
	再考虑在第 2 个位置切割，aa | bb
	这样我们只需要知道 bb 的所有结果，然后加 1，记为 m2


	aabb
	再考虑在第 3 个位置切割，aab|b
	因为 aab 不是回文串，所有直接跳过

	aabb
	再考虑在第 4 个位置切割，aabb |
	因为 aabb 不是回文串，所有直接跳过

	此时只需要比较 m1 和 m2 的大小，选一个较小的即可。
*/
var dp [][]bool // dp 把每个子串是否是回文串，提前存起来 😅😅😅
var str string
var memo map[int]int // 缓存重复解的计算

func minCut(s string) int {
	str = s
	sLen := len(str)
	dp = make([][]bool, sLen)
	memo = make(map[int]int, 0)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}
	//
	for i := 1; i <= sLen; i++ {
		for j := 0; j <= sLen-i; j++ {
			// 没看懂😅😅😅
			cut := i + j - 1
			if str[j] == str[cut] && (i < 3 || dp[j+1][cut-1]) {
				dp[j][cut] = true
			}
		}
	}
	return helper(0)
}

func helper(start int) int {
	if val, ok := memo[start]; ok {
		return val
	}
	// base case，递归出口。长度是1，最小切割次数0
	if dp[start][len(str)-1] {
		return 0
	}

	min := (1 << 32)
	for i := start; i < len(str); i++ {
		// 只考虑回文串，
		if dp[start][i] {
			// 和之前的值比较选一个较小的
			min = MinInt(min, 1+helper(i+1))
		}
	}
	memo[start] = min
	return min
}

/*
	方法2：动态规划
	question 😅😅😅
	// state: dp[i] "前i"个字符组成的子字符串需要最少几次cut(个数-1为索引)
	// function: dp[i] = MIN{dp[j]+1}, j < i && [j+1 ~ i]这一段是一个回文串
	// intialize: dp[i] = i - 1 (dp[0] = -1)
	// answer: dp[s.length()]
*/
func minCut(s string) int {
	sLen := len(s)
	if sLen == 0 || sLen == 1 {
		return 0
	}

	dp := make([]int, sLen+1)
	dp[0] = -1
	dp[1] = 0
	for i := 1; i <= sLen; i++ {
		dp[i] = i - 1
		for j := 0; j < i; j++ {
			if isPalindrome(s, j, i-1) {
				dp[i] = MinInt(dp[i], dp[j]+1)
			}
		}
	}
	return dp[sLen]
}

// 判断是否是回文串
func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}