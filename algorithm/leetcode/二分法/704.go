/*
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	你可以假设 nums 中的所有元素是不重复的。
	n 将在 [1, 10000]之间。
	nums 的每个元素都将在 [-9999, 9999]之间。
	0103 check写错了
*/

// 模板1：满足某种情况的最小的元素。是否存在一个目标值
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 区间分为[l, mid]和[mid + 1, r]两部分 😅😅😅
	for left < right {
		// 计算mid时不需要加1 😅😅😅😅
		mid := (left + right) >> 1
		// check条件 😅😅😅
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}