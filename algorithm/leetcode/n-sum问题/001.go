/*
	给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
*/

/*
	方法1： 两次循环
	方法2：一次循环+hash
*/
func twoSum1(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		// 判断hash中是否存在target-nums[i]
		if _, ok := hash[target-nums[i]]; ok {
			// 直接返回结果
			return []int{hash[target-nums[i]], i}
		}
		// nums[i]存入hash中作为键，下标i作为value
		hash[nums[i]] = i
	}
	// 未找到符合条件的情况
	return []int{-1, -1}
}

