/*
	13. 找出数组中重复的数字

	给定一个长度为 n 的整数数组 nums，数组中所有的数字都在 0∼n−1 的范围内。
	数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
	请找出数组中任意一个重复的数字。
	注意：如果某些数字不在 0∼n−1 的范围内，或数组中不包含重复数字，则返回 -1；

	给定 nums = [2, 3, 5, 4, 3, 2, 6, 7]。返回2或3
*/

func duplicateInArray(nums []int) int {
	hash := make(map[int]int)
	// 不在范围内的数字
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > len(nums)-1 {
			return -1
		}
	}
	// 利用hash找出重复数字
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] += 1
		if hash[nums[i]] > 1 {
			return nums[i]
		}
	}
	return -1
}