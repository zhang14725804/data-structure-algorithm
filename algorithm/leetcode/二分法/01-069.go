// 精确查找
func mySqrt(x int) int {
	// 😅😅 [left, right]两端都闭的区间
	left, right := 0, x
	// 1. 😅😅 循环条件 【<=】
	for left <= right {
		mid := (right + left) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		}
	}
	// 4. 取right，而不是left
	return right
}

/*
	模板1：【满足某种情况的最小的元素】；[l, r]区间划分为[l, mid] 和 [mid+1, r]
*/
func mySqrt(x int) int {
	// （1） 0，1 特殊处理
	if x == 0 || x == 1 {
		return x
	}
	left, right := 0, x
	for left < right {
		mid := (left + right) >> 1
		// 注意check条件 😅
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 😅😅😅 question 为什么返回left-1
	return left - 1
}