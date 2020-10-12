/*
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

	你算法的时间复杂度应该为 O(n2) 。
	进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

*/

/*
	动态规划思路：
	（1）状态表示。集合（所有以i结尾的上升子序列），属性：长度最大值
	（2）状态计算。
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	// 表示以第 i 个数字为结尾的最长上升子序列的长度
	dp := make([]int, n)
	var res int
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 如果前边的某个数 nums[j] < nums[i] ，那么我们可以将第 i 个数接到第 j 个数字的后边作为一个新的上升子序列
			if nums[j] < nums[i] {
				// 这里不好理解（todo）
				dp[i] = MaxInt(dp[i], dp[j]+1)
			}
		}
		res = MaxInt(res, dp[i])
	}
	return res
}

/*
	方法二不理解（todo）

	dp[i] 表示长度为 i + 1 的所有上升子序列的末尾的最小值
*/