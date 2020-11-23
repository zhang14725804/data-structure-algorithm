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
				dp[j] = dp[j] + nums[j+i]
				if i != 0 && dp[j] == 0 {
					return true
				}
			}
		}
		return false
	}

	// 当i=k时，dp[j]表示以j为起始下标，nums中连续k+1个整数的和
	// 如当i=0时，相当于将nums拷贝到dp
	// i=1时，dp[0]相当于以0为起始下标，nums中2个整数的和，即nums[0]+nums[1]
	// 每次计算时都可以用原来的dp进行更新，而不用一个个去加

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			dp[j] = (dp[j] + nums[j+i]) % k
			if i != 0 && dp[j] == 0 {
				return true
			}
		}
	}
	return false
}