/*
	已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
	在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
	例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
	如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

	question: 重复元素到底影响什么😅😅😅
*/
func search(nums []int, target int) bool {
	n := len(nums) - 1
	if n < 0 {
		return false
	}
	left, R := 0, n
	// 😅😅😅😅😅😅去除首尾的重复元素
	for nums[left] == nums[R] && left < R {
		R--
	}
	// 注意这里。😅😅😅
	right := R
	// 用模板2，【满足某种情况的最大的元素】；找到最大值(分界点)
	for left < right {
		mid := (left + right + 1) >> 1
		// 😅😅 注意check条件【>=】
		if nums[mid] >= nums[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 判断target在那一边
	if target >= nums[0] {
		left = 0
	} else {
		left = right + 1
		right = R
	}

	// // 用模板1，【满足某种情况的最小的元素】，精确查找
	for left < right {
		mid := (left + right) >> 1
		// 😅😅 注意check条件【>=】
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right] == target
}