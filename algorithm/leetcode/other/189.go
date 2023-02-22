/*
	给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

	尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
	要求使用空间复杂度为 O(1) 的 原地 算法。
*/
// 方法1：三次旋转法
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}