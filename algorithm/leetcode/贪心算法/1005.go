/*
	贪心的思路，局部最优：让绝对值大的负数变为正数，当前数值达到最大，整体最优：整个数组和达到最大。

	那么如果将负数都转变为正数了，K依然大于0，此时的问题是一个有序正整数序列，如何转变K次正负，让 数组和 达到最大。

	那么又是一个贪心：局部最优：只找数值最小的正整数进行反转，当前数值可以达到最大（例如正整数数组{5, 3, 1}，反转1 得到-1 比 反转5得到的-5 大多了），全局最优：整个 数组和 达到最大。

	（1）将数组按照绝对值大小从大到小排序，注意要按照绝对值的大小
	（2）从前向后遍历，遇到负数将其变为正数，同时K--
	（3）如果K还大于0，那么反复转变数值最小的元素，将K用完
	（4）求和
*/
func largestSumAfterKNegations(nums []int, k int) int {
	// 第一步
	nums = BubbleSort(nums)
	// 第二部
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	// 第三步
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	// 第四步
	res := 0
	for v, _ := range nums {
		res += v
	}
	return res
}