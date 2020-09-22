/*
	在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

	todo：快排，小顶堆
*/

// 方法1：直接采用快排对数组进行排序
func findKthLargest(nums []int, k int) int {
	nums = quickSort(nums)
	return nums[len(nums)-k]
}

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
	// 将数据分为左右两部分
	for j := start; j < end; j++ {
		if nums[j] > pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	// todo：这什么意思
	nums[end], nums[i] = nums[i], nums[end]

	c := i - start + 1
	if c == k {
		return nums[i]
	} else if c < k {
		// 右边
		return helper(i+1, end, k-c)
	} else {
		// 左边
		return helper(start, i-1, k)
	}
}