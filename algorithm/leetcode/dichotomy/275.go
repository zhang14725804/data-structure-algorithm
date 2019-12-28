package leetcode

// Leetcode275  H-Index II  (二分法)
func Leetcode275(nums []int) int {
	// 这里right大小就是len(nums)，不需要减1
	left, right := 0, len(nums)
	for left < right {
		// 注意临界条件
		mid := (left + right + 1) >> 1
		if nums[len(nums)-mid] >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}
