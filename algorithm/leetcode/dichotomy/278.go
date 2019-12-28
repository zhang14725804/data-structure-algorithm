package leetcode

// 大概就这个意思
func isBadVersion(version int) bool {
	return true
}

// Leetcode278 （二分法）
func Leetcode278(n int) int {
	// 应用模板1
	left, right := 1, n
	for left < right {
		mid := (left + right) >> 1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
