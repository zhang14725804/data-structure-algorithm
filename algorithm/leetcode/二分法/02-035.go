/*
	方法1：找到大于等于target得最左边的数（模板1）
*/
func searchInsert(nums []int, target int) int {
	// 边界条件
	if len(nums) == 0 || nums[len(nums)-1] < target {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	// 使用模板1，找................vooooooooo中的v，这种情况😅😅😅
	for l < r {
		mid := (l + r) >> 1
		// [l, r]区间划分为[l, mid] 和 [mid+1, r]
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
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
	// question 最后需要判断下。那模板1为什么不需要判断呢 😅😅😅😅😅😅😅😅😅
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