/*
	给定一个包含 非负数 的数组和一个目标 整数 k，编写一个函数来判断该数组是否含有连续的子数组，其大小至少为 2，且总和为 k 的倍数，即总和为 n*k，其中 n 也是一个整数。
	数组的长度不会超过 10,000 。

	懵(question),都不懂
	方法1：前缀和与hash
	方法2：动态规划
*/
func checkSubarraySum(nums []int, k int) bool {
	dp := make([]int, 10010)
	n := len(nums)
	if n < 2 {
		return false
	}
	if k == 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n-i; j++ {
				dp[j] = dp[j] + nums[i+j]
				if i != 0 && dp[j] == 0 {
					return true
				}
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			dp[j] = (dp[j] + nums[i+j]) % k
			if i != 0 && dp[j] == 0 {
				return true
			}
		}
	}
	return false
}