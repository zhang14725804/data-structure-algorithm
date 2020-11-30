/*
	峰值元素是指其值大于左右相邻值的元素。
	给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
	数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
	你可以假设 nums[-1] = nums[n] = -∞。

	你的解法应该是 O(logN) 时间复杂度的。
*/
func findPeakElement1(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 模板1；[l, r]区间划分为[l, mid] 和 [mid+1, r]
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func findPeakElement2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		// 模板2；[l, r]区间划分为[l, mid - 1] 和 [mid, r]
		mid := (l + r + 1) >> 1
		if nums[mid-1] < nums[mid] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}