/*
	给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。
	对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

	返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

	数组非空，且长度不会超过 20 。
	初始的数组的和不会超过 1000 。
	保证返回的最终结果能被 32 位整数存下。

*/
var ans int
var nums []int
var target int

// 方法1：回溯算法
func findTargetSumWays1(_nums []int, S int) int {
	target = S
	nums = _nums
	backtrack(0)
	return ans
}
func backtrack(i int) {
	if i == len(nums) {
		if target == 0 {
			ans++
		}
		return
	}
	// 选择加号
	target += nums[i]
	backtrack(i + 1)
	target -= nums[i]
	// 选择减号
	target -= nums[i]
	backtrack(i + 1)
	target += nums[i]

}

// 方法2：递归+cache
func findTargetSumWays(nums []int, target int) int {

}

// 方法3：动态规划，背包问题(question)，不懂（😅）