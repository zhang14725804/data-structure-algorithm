/*
	给你一个整数数组 nums ，和一个整数 target 。
	该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。
	请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。


*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 找到最小值
	l, r := 0, len(nums)-1
	for l < r {
		// 模板1：满足某种情况的最小的元素。是否存在一个目标值
		mid := (l + r) >> 1
		if nums[mid] <= nums[len(nums)-1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 判断目标值范围
	if target <= nums[len(nums)-1] {
		// 在右边
		r = len(nums) - 1
	} else {
		// 在左边
		l = 0
		r--
	}

	// 寻找目标值
	for l < r {
		// 模板1：满足某种情况的最小的元素。是否存在一个目标值
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	// 注意:这里只能用l
	if nums[l] == target {
		return l
	}
	return -1
}