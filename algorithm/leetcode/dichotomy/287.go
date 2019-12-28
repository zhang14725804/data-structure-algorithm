package leetcode

// Leetcode287  Find the Duplicate Number (二分法)
func Leetcode287() int {
	nums := []int{1, 3, 4, 2, 2}
	/*
		左右两边不可能都 <=，至少有一边大于
		每次把答案的范围缩小
	*/
	n := len(nums)
	left, right := 1, n-1
	for left < right {
		mid := (left + right) >> 1
		cnt := 0
		// 判断难理解
		for _, x := range nums {
			if x >= left && x <= mid {
				cnt++
			}
		}
		if cnt > mid-left+1 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
