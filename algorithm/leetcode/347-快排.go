/*
	给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

	你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
	你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
	题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
	你可以按任意顺序返回答案。
*/

// 方法2：快速排序变形（🔥🔥🔥不需要对整个数组排序，切分即可）  有问题！！！！
func getLeastNumbers(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}
	// ⚠️注意最后一个参数传入我们要找的下标（第k小的数下标是k-1）
	return quickSearch(nums, 0, len(nums)-1, k-1)
}

func quickSearch(nums []int, low, high, k int) []int {
	// 每快排切分1次，找到排序后下标为j的元素，如果j恰好等于k就返回j以及j左边所有的数；
	j := partition(nums, low, high)
	res := make([]int, j+1)
	fmt.Println(nums)
	if j == k {
		copy(res, nums)
		return res
	}
	// 否则根据下标j与k的大小关系来决定继续切分左段还是右段。
	if j > k {
		quickSearch(nums, low, j-1, k)
	}
	return quickSearch(nums, j+1, high, k)
}

func partition(nums []int, low, high int) int {
	fmt.Println(low, high)
	v := nums[low]
	i := low
	j := high + 1
	for {
		for i+1 <= high && nums[i+1] < v {
			i++
		}
		for j-1 >= low && nums[j-1] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[low] = nums[j]
	nums[j] = v
	return j
}