package leetcode

// Leetcode033 Search in Rotated Sorted Array（二分法）
func Leetcode033(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	// 找打最小值
	left, right, last := 0, n-1, n-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] <= nums[last] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 不好理解，在153题上再循环一次
	if target <= nums[last] {
		right = n - 1
	} else {
		left = 0
		right--
	}
	for left < right {
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
