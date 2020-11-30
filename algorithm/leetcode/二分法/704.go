/*
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

	你可以假设 nums 中的所有元素是不重复的。
	n 将在 [1, 10000]之间。
	nums 的每个元素都将在 [-9999, 9999]之间。

	question:
		(1) 循环条件为什么是left <= right
		left <= right终止区间[right+1,right]
		left < right终止区间[right,right],漏掉了一个数字
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// 模板1：满足某种情况的最小的元素。是否存在一个目标值
		mid := (left + right) >> 1
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

func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 循环条件：开区间
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 跳过mid
			left = mid + 1
		} else if nums[mid] > target { // 跳过mid
			right = mid - 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 循环条件：闭区间
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { // 跳过mid
			left = mid + 1
		} else if nums[mid] > target { // 跳过mid
			right = mid - 1
		}
	}
	return -1
}