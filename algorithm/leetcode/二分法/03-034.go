/*
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置

	question  😅😅😅😅😅😅
*/
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}

	left, right := 0, n-1
	// 使用模板1，找................vooooooooo中的v，这种情况😅😅😅
	for left < right {
		// 区间分为[l, mid]和[mid + 1, r]两部分 😅😅😅
		mid := (left + right) >> 1
		// check函数
		if nums[mid] >= target {
			right = mid // 😅😅😅满足条件，移动有边界
		} else {
			left = mid + 1 // 😅😅😅不满足，移动左边界
		}
	}
	// 判断并找到左端点 😅😅😅
	if nums[left] != target {
		return []int{-1, -1}
	}
	start := left

	left, right = 0, n-1
	// 使用模板2，找oooooooov...................中的v，这种情况😅😅😅
	for left < right {
		// 区间分为[l, mid-1]和[mid, r]两部分 😅😅😅，此时为了防止死循环，计算mid时需要加1。
		mid := (left + right + 1) >> 1
		// check函数
		if nums[mid] <= target {
			left = mid // 😅😅😅满足条件，移动左边界
		} else {
			right = mid - 1 // 😅😅😅不满足条件，移动有边界
		}
	}
	// 打完收工
	end := left
	return []int{start, end}
}