/*
	给定一个未排序的整数数组，找出最长连续序列的长度。
	要求算法的时间复杂度为 O(n)。
*/

// 思路，先排序，然后找
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	res := 0
	j := 1
	for i := 0; i < len(nums); i++ {
		// 去重
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		if i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			j++
		} else {
			res = compare(res, j, true)
			j = 1
		}
	}
	return res
}