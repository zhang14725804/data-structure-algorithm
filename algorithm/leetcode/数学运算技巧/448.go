/*
	集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
	给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

	给定数组的长度范围是 [2, 10000]。
	给定的数组是无序的。

	(question)不懂(😅)
	ps：关键点在于元素和索引是成对儿出现的，常用的方法是排序、异或、映射。
*/
func findErrorNums(nums []int) []int {
	n := len(nums)
	dup := -1
	for i := 0; i < n; i++ {
		// 现在的元素是从 1 开始的
		index := abs(nums[i]) - 1
		// nums[index] 小于 0 则说明重复访问
		if nums[index] < 0 {
			dup = abs(nums[i])
		} else {
			nums[index] *= -1
		}
	}

	missing := -1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			// 将索引转换成元素
			missing = i + 1
		}
	}
	return []int{dup, missing}
}