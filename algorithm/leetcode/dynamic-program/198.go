package leetcode

import (
	"data-structure-algorithm/algorithm/common"
	"fmt"
)

// Leetcode198 “House Robber”
func Leetcode198() int {
	/*
		f[i]表示在前i个数中选，所有不选nums[i]的选法最大值
		g[i]表示在前i个数中选，所有选择nums[i]的选法最大值

		f[i] = max(f[i-1],g[i-1])
		g[i] = f[i-1]+nums[i]
	*/
	nums := []int{1, 2, 3, 1}
	n := len(nums)
	f := make([]int, n+1)
	g := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = common.Max(f[i-1], g[i-1])
		g[i] = f[i-1] + nums[i-1]
	}
	fmt.Println(f)
	fmt.Println(g)
	return common.Max(f[n], g[n])
}
