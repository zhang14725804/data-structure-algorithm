/*
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。
	( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
	请找出其中最小的元素。
	注意数组中可能存在重复的元素。

*/
func findMin1(nums []int) int {
	start, end := 0, len(nums)-1
	// 预处理，保证所有重复数字不在两段里出现，保证切割的位置不要是重复数字 （question，没懂😅）
	for nums[start] == nums[end] && start < end {
		end--
	}
	for start < end {
		// 是否存在一个目标值（精确查找）
		mid := (start + end) >> 1
		if nums[mid] > nums[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return nums[start]
}

func findMin(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := (start + end) >> 1
		if nums[mid] > nums[end] {
			start = mid + 1
		} else if nums[mid] < nums[end] {
			end = mid
		} else {
			//
			end--
		}
	}
	return nums[start]
}