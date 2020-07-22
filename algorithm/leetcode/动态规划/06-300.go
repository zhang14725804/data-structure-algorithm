import (
	"data-structure-algorithm/algorithm/common"
)

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
	（2）状态计算。todo，状态计算蒙圈了，不懂

*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			// 倒数第二个小于最后一个
			if nums[j] < nums[i] {
				f[i] = common.Max(f[i], f[j]+1)
			}
		}
	}
	var res int
	for k := 0; k < n; k++ {
		res = common.Max(res, f[k])
	}
	return res
}
