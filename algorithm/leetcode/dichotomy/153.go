package leetcode

// Leetcode153 Find Minimum in Rotated Sorted Array（二分法）
func Leetcode153(nums []int) int {
	left, right, last := 0, len(nums)-1, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		// 判断条件:所有右边的元素都小于最后一个元素，所有左边的元素都大于最后一个元素
		if nums[mid] <= nums[last] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
