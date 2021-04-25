/*
	求平方根
*/
// 模板2：满足某种情况的最大的元素；[l, r]区间划分为[l, mid - 1] 和 [mid, r]
func mySqrt1(x int) int {
	l, r := 0, x
	for l < r {
		// 求mid的时候需要上取整
		mid := (l + r + 1) >> 1
		if mid*mid <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

// 模板1；[l, r]区间划分为[l, mid] 和 [mid+1, r]
func mySqrt2(x int) int {
	l, r := 0, x
	for l < r {
		mid := (l + r) >> 1
		if mid*mid > x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// (question)这里为什么减1，x是零的情况还需要特殊处理
	return l - 1
}