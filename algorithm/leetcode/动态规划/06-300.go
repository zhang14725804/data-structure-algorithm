/*
	给定一个无序的整数数组，找到其中最长上升子序列的长度。

	示例:

	输入: [10,9,2,5,3,7,101,18]
	输出: 4
	解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

	说明:

	可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
	你算法的时间复杂度应该为 O(n2) 。
	进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

	动态规划思路：
	（1）状态表示。集合（所有以i结尾的上升子序列），属性：长度最大值
	（2）状态计算。

	https://leetcode.wang/leetcode-300-Longest-Increasing-Subsequence.html

*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)
	var res int
	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			// 如果前边的某个数 nums[j] < nums[i] ，那么我们可以将第 i 个数接到第 j 个数字的后边作为一个新的上升子序列
			if nums[j] < nums[i] {
				f[i] = compare(f[i], f[j]+1, true)
			}
		}
		res = compare(res, f[i], true)
	}
	return res
}