/*
	1. 声明一个map，【key当前元素】，【value当前元素索引值】
	2. 遍历数组，检查是否存在target-nums[i]的元素
*/
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		// 判断hash中是否存在target-nums[i]
		key := target - nums[i]
		if idx, ok := hash[key]; ok {
			// 直接返回结果
			return []int{idx, i}
		}
		// nums[i]存入hash中作为键，下标i作为value
		hash[nums[i]] = i
	}
	// 未找到符合条件的情况
	return []int{-1, -1}
}


