/*
	给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

*/

// 暴力枚举
func threeSumClosest2(nums []int, target int) int {
	sub := INT_MAX
	sum := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if absInt(nums[i]+nums[j]+nums[k]-target) < sub {
					sum = nums[i] + nums[j] + nums[k]
					sub = absInt(sum - target)
				}
			}
		}
	}
	return sum
}

/*

	排序和双指针

	在数组 nums 中，进行遍历，每遍历一个值利用其下标i，形成一个固定值 nums[i]
	再使用前指针指向 start = i + 1 处，后指针指向 end = nums.length - 1 处，也就是结尾处
	根据 sum = nums[i] + nums[start] + nums[end] 的结果，判断 sum 与目标 target 的距离，如果更近则更新结果 ans
	同时判断 sum 与 target 的大小关系，因为数组有序，如果 sum > target 则 end--，如果 sum < target 则 start++，如果 sum == target 则说明距离为 0 直接返回结果

*/
func threeSumClosest(nums []int, target int) int {
	nums = quickSort(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if absInt(target-sum) < absInt(target-ans) {
				ans = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return ans
			}
		}
	}
	return ans
}