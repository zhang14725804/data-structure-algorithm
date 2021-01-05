/*
	集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
	给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
*/
// 先遍历一次数组，找出重复出现的数字，在遍历一次used数组，找出缺失的
func findErrorNums1(nums []int) []int {
	n := len(nums)
	used := make([]bool, n)
	dup := -1
	for i := 0; i < n; i++ {
		if !used[nums[i]-1] {
			used[nums[i]-1] = true
		} else {
			dup = nums[i]
		}

	}
	missing := -1
	for i := 0; i < n; i++ {
		if !used[i] {
			missing = i + 1
		}
	}
	return []int{dup, missing}
}

//
func findErrorNums(nums []int) []int {
	n := len(nums)
	dup := -1
	for i := 0; i < n; i++ {
		// question 如何不使用额外空间判断某个索引有多少个元素对应呢？
		// 通过将每个索引对应的元素变成负数，以表示这个索引被对应过一次了
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			dup = abs(nums[i])
		} else {
			nums[index] *= -1
		}
	}
	missing := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			// // 将索引转换成元素
			missing = i + 1
		}
	}
	return []int{dup, missing}
}