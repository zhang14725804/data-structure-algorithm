/*
	给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
	说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

	todo：和169类似（摩尔投票法）
*/

// 方法1：hash
func majorityElement(nums []int) []int {
	n := len(nums)
	hash := make(map[int]int, 0)
	res := make([]int, 0)
	// todo：代码有问题
	for i := 0; i < n; i++ {
		hash[nums[i]]++
		if hash[nums[i]] == 2*n/3 || len(res) == 2 {
			return res
		}
		if hash[nums[i]] > n/3 {
			res = append(res, nums[i])
		}
	}
	return res
}