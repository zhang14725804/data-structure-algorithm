

/*
	给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

*/
func maxProduct(nums []int) int {
	maxArr := make([]int, 0)
	minArr := make([]int, 0)
	maxArr = append(maxArr, nums[0])
	minArr = append(minArr, nums[0])

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// todo：注意这里
		maxArr = append(maxArr, compare(compare(nums[i], maxArr[i-1]*nums[i], true), minArr[i-1]*nums[i], true))
		minArr = append(minArr, compare(compare(nums[i], maxArr[i-1]*nums[i], false), minArr[i-1]*nums[i], false))
		max = compare(max, maxArr[i], true)
	}
	return max
}
func compare(a, b int, max bool) int {
	// max 是否返回最大值
	if a > b {
		if max == true {
			return a
		}
		return b
	}
	if max == true {
		return b
	}
	return a
}
