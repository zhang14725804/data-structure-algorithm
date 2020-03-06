/*
	26. Remove Duplicates from Sorted Array

	双指针算法（很巧妙）

	todos:有问题
*/
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	k:=1
	for j:=1;j<len(nums);j++{
		if nums[j]!=nums[j-1]{
			k+=1
			nums[k] = nums[j]
		}
	}
	return k
}