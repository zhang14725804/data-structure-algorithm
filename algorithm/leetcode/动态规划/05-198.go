package leetcode

import (
	"data-structure-algorithm/algorithm/common"
)

/*
	你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
	给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

	动态规划思路：

	f[i]表示在前i个数中选，所有不选nums[i]的选法最大值
	g[i]表示在前i个数中选，所有选择nums[i]的选法最大值

	f[i] = max(f[i-1],g[i-1])
	g[i] = f[i-1]+nums[i]
*/
func rob(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	g := make([]int, n+1)
	// 从1开始，避免处理边界问题
	for i := 1; i <= n; i++ {
		f[i] = common.Max(f[i-1], g[i-1])
		// 数组下标从0开始
		g[i] = f[i-1] + nums[i-1]
	}
	return common.Max(f[n], g[n])
}
