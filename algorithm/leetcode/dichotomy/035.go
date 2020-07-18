/*
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
	如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

	你可以假设数组中无重复元素。
*/
func searchInsert(nums []int, target int) int {
	length := len(nums)
	// 方法二：二分法（临界条件不好判断）
	if length == 0 || nums[length-1] < target {
		return length
	}
	l, r := 0, length-1
	for l < r {
		// 这种通常题目描述为满足某种情况的最小的元素。是否存在一个目标值（精确查找）
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
func searchInsert(nums []int, target int) int {
	length := len(nums)
	// 方法一 普通循环
	for i := 0; i < length; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return length
}
