/*
	给你一个整数数组 nums ，和一个整数 target 。
	该整数数组原本是按升序排列，但输入时在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。
	请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
*/
func search(nums []int, target int) int {
	n := len(nums) - 1
	if n <= 0 {
		return -1
	}
	// 用模板1，找到最小值所在的位置
	left, right := 0, n
	for left < right {
		mid := (left + right) >> 1
		// 😅😅 注意check条件
		if nums[mid] <= nums[n] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// 判断目标值范围在最小值左边还是右边, 判断条件是 😅【<=】
	if target <= nums[n] {
		// 在右边
		right = n
	} else {
		// 在左边
		left = 0
		right--
	}

	// 采用模板1，查找精确值
	for left < right {
		mid := (left + right) >> 1
		// 😅😅 注意check条件
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	// 注意:这里只能用l
	if nums[left] == target {
		return left
	}
	return -1
}
