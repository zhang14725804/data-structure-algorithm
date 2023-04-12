// 寻找左侧边界的二分搜索
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 😅😅 [left, right) 左闭右开
	left, right := 0, len(nums)
	// 1. 😅😅 循环条件 【<】
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return left
}

/*
	方法2：找到小于target得最右边的数（模板2）
*/
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := (right + left + 1) >> 1
		if nums[mid] < target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// question 最后需要判断下。模板1为什么不需要判断呢 😅😅😅😅😅😅😅😅😅
	if nums[right] < target {
		return right + 1
	}
	return right
}

/*
	这种写法对理解69题方法2有帮助；模板1的另一种写法
	question 这他么什么意思 😅😅😅
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || nums[len(nums)-1] < target {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}