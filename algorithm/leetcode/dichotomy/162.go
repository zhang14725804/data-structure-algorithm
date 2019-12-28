package leetcode

// Leetcode162  Find Peak Element（相邻两个值不同，可以二分）有难度 (二分法)
func Leetcode162(nums []int) int {
	/*
		出现拐点，这个点就是峰值。不出现拐点，边界就是峰值
	*/
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
