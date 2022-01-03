/*
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

	你算法的时间复杂度应该为 O(n2) 。
	进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

/*
	方法1：动态规划
	😅😅😅😅😅😅
	动态规划思路：
		（1）状态表示。集合（所有以i结尾的上升子序列），属性：长度最大值
		（2）状态计算。
	0103 没懂dp的意思
*/
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	// dp[i]：表示以第 i 个数字为结尾的最长上升子序列的长度
	dp := make([]int, n)
	var res int
	for i := 0; i < n; i++ {
		// dp[i] 所有元素置 1，含义是每个元素都至少可以单独成为子序列，此时长度都为 1。
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
	方法二：动态规划+二分查找
	( question ) 😅😅😅 没懂，难懂
*/
func lengthOfLIS(nums []int) int {
	res := 0 // tails当前长度
	n := len(nums)
	// tails[k] 的值代表 长度为 k+1k+1 子序列 的尾部元素值
	tail := make([]int, n)
	for i := 0; i < n; i++ {
		poker := nums[i]
		left, right := 0, res
		for left < right {
			mid := (left + right) >> 1
			if tail[mid] >= poker {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == res {
			res++
		}
		tail[left] = poker
	}
	return res
}