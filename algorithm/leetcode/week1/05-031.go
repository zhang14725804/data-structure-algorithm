/*
	实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
	如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

	思路（有点脑筋急转弯的意思）：从后向前找到第一个不是从小到大排列的数，将这个数和后边第一个比它大的数交换位置，再把这个数后面的数反转
	例如：8454321 --> 8544321 --> 8512344
*/
func nextPermutation(nums []int) {
	// 从最后一个数字开始遍历数组
	for i := len(nums) - 1; i > 0; i-- {
		// 存在后一个数比前一个数大 123--> 132
		if nums[i-1] < nums[i] {
			j := i
			for j+1 < len(nums) && nums[j+1] > nums[i-1] {
				j++
			}
			// 这里不好懂
			nums[i-1], nums[j] = nums[j], nums[i-1]
			reverse(nums, i, len(nums)-1)
			return
		}
	}
	reverse(nums, 0, len(nums)-1)
}

func reverse(a []int, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}