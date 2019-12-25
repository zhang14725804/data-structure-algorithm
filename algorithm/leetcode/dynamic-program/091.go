package leetcode

import (
	"fmt"
	"strconv"
)

// Leetcode091 动态规划
func Leetcode091() int {
	s := "226"
	if len(s) == 0 {
		return 0
	}
	n := len(s)
	// 声明数组
	dp := make([]int, n+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= n; i++ {
		//
		last, _ := strconv.Atoi(s[i-1 : i])
		if last >= 1 && last <= 9 {
			dp[i] += dp[i-1]
		}
		last, _ = strconv.Atoi(s[i-2 : i])
		if last >= 10 && last <= 26 {
			dp[i] += dp[i-2]
		}
	}
	fmt.Println(dp)
	return dp[n]
}
