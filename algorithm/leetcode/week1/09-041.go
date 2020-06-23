/*
	给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
	要求：你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

	todo：思路清奇啊
*/
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 👉：精华
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i]{
			nums[i],nums[nums[i]-1] =  nums[nums[i]-1],nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1{
			return i+1
		}
	}
	return len(nums)+1
}