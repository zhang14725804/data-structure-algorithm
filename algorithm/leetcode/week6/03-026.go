/*
	Remove Duplicates from Sorted Array

	双指针算法（很巧妙）

	todos:有问题
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// k存所有不重复的数
	k := 1
	// j从前向后遍历
	for j := 1; j < len(nums); j++ {
		// 判断是否和上一个数相同
		if nums[j] != nums[j-1] {
			// nums[k++] = nums[j]用下边两行代替
			nums[k] = nums[j]
			k++
		}
	}
	return k
}