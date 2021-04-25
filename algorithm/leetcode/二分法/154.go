/*
	假设按照升序排序的数组在预先未知的某个点上进行了旋转。
	( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
	请找出其中最小的元素。
	注意数组中可能存在【重复】的元素。😅😅😅😅😅😅😅😅😅
*/
//  精确查找，模板1
func findMin(nums []int) int {
	n := len(nums) - 1
	left, right := 0, n
	// 👉： question  预处理，【循环删除】重复元素。保证所有重复数字不在两段里出现，保证切割的位置不要是重复数字
	// 把for写成了if，居然不会调试，😅😅😅
	for nums[left] == nums[right] && left < right {
		right--
	}
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
		// check条件也可以写成这样
		// if nums[mid] > nums[right] {
		// 	left = mid + 1
		// } else {
		// 	right = mid
		// }
	}
	return nums[left]
}