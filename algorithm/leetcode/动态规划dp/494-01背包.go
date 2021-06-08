/*
	给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。
	对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

	返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

	数组非空，且长度不会超过 20 。
	初始的数组的和不会超过 1000 。
	保证返回的最终结果能被 32 位整数存下。

*/

/*
	方法1：回溯算法
	question 😅😅😅 这道题为什么不需要for循环 😅😅😅
*/
var ans int
var nums []int
var target int

func findTargetSumWays(N []int, T int) int {
	target = T
	nums = N
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

/*
	方法3：动态规划，背包问题
	如何转化为01背包问题呢。 （😅😅😅）
*/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum {
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	// 此时问题就转化为，装满容量为x背包，有几种方法
	bgSize := (target + sum) / 2
	// dp[j] 表示：填满j（包括j）这么大容积的包，有dp[i]种方法
	dp := make([]int, bgSize+1)
	// 初始化。dp[0] = 1，理论上也很好解释，装满容量为0的背包，有1种方法，就是装0件物品
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := bgSize; j >= nums[i]; j-- {
			// 😅😅😅 递归公式，求装满背包有几种方法的情况😅😅😅
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bgSize]
}