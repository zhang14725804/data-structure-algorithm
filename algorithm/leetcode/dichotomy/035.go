package leetcode

// Leetcode035 Search Insert Position（常规方法和二分法）
func Leetcode035() int {
	nums := []int{1, 3, 5, 6}
	target := 5
	length := len(nums)
	// 方法一 普通循环
	// for i := 0; i < length; i++ {
	// 	if nums[i] >= target {
	// 		return i
	// 	}
	// }
	// return length

	// 方法二：二分法（临界条件不好判断）
	if length == 0 || nums[length-1] < target {
		return length
	}
	l, r := 0, length-1
	for l < r {
		// 这里注意括号
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
