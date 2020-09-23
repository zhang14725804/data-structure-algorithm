/*
	给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

	这题意什么意思😅：是否有相同的数字，两个数字最多相隔k
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if index, ok := hash[nums[i]]; ok {
			if i-index <= k {
				return true
			}
		}

		hash[nums[i]] = i
	}
	return false
}