/*
	寻找旋转排序数组中的最小值.假设按照升序排序的数组在预先未知的某个点上进行了旋转。
*/

// 精确查找
func findMin(nums []int) int {
	left, right, n := 0, len(nums)-1, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		// 😅😅😅😅 question check条件
		if nums[mid] <= nums[n] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}