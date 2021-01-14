/*
	给定一个整数数组，判断是否存在重复元素。
	如果任意一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。

	hash，或者排序
*/

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = 1
	}
	return false
}