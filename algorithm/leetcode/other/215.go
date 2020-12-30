/*
	在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
	排序时间复杂度是 O(NlogN)

	两种方法：二叉堆（优先队列）、快速选择算法
*/

/*
	方法2：部分数据快排

	随机选择一个分区点，左边都是大于分区点的数，右边都是小于分区点的数。左部分的个数记做 m。

	如果 k == m + 1，我们把分区点返回即可。

	如果 k > m + 1，说明第 k 大数在右边，我们在右边去寻找第 k - m - 1 大的数即可。

	如果 k < m + 1，说明第 k 大数在左边，我们在左边去寻找第 k 大的数即可。
*/

var nums []int

func findKthLargest(_nums []int, k int) int {
	nums = _nums
	return helper(0, len(nums)-1, k)
}

func helper(start, end, k int) int {
	i := start
	pivot := nums[end] // 基准点
	for j := start; j < end; j++ {
		if nums[j] > pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	// ( question ) 将 pivot 值交换到正确的位置
	nums[end], nums[i] = nums[i], nums[end]

	c := i - start + 1
	if c == k {
		return nums[i]
	} else if c < k {
		return helper(i+1, end, k-c)
	} else {
		return helper(start, i-1, k)
	}
}