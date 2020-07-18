/*
	定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

	你的算法时间复杂度必须是 O(log n) 级别。

	如果数组中不存在目标值，返回 [-1, -1]。

	思路1：借助hash
	思路2：二分法
*/

func searchRange1(nums []int, target int) []int {
	n := len(nums)
	// 常规map方法
	mp := make(map[int][]int)
	for i := 0; i < n; i++ {
		if arr, ok := mp[nums[i]]; ok {
			arr = append(arr, i)
			mp[nums[i]] = arr
		} else {
			mp[nums[i]] = []int{i}
		}
	}
	res, _ := mp[target]
	// 判空
	if len(res) == 0 {
		return []int{-1, -1}
	}
	l := res[0]
	r := res[len(res)-1]
	return []int{l, r}
}

// 二分法（两次二分）
func searchRangenums []int, target int) []int{
	n := len(nums)
	// 判空
	if n == 0 {
		return []int{-1, -1}
	}
	
	// 左边界
	left, right := 0, n-1
	for left < right {
		// 这种通常题目描述为满足某种情况的最小的元素。是否存在一个目标值（精确查找）
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[right] != target {
		return []int{-1, -1}
	}
	start := right

	// 右边界
	left, right = 0, n-1
	for left < right {
		// 这种通常题目描述为满足某种情况的最大的元素
		mid := (left + right + 1) >> 1
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	end := right
	return []int{start, end}
}