/*
	给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。
	要求！！！：你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

	思路1：排序，然后再找
	ps：优先考虑时间复杂度，在考虑空间复杂度
	思路2：hash表（空间复杂度有问题😅）
*/
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 👉：精华
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// hash方法
func firstMissingPositive(nums []int) int {
	hash := make(map[int]int)
	// （1）依次存入hash
	for _, n := range nums {
		hash[n] = 1
	}
	// （2）从1开始判断是否存在
	res := 1
	for hash[res] > 0 {
		res++
	}
	// 返回结果
	return res
}