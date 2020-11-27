/*
	给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置
	需要两次二分

	二分法套路1
*/
func searchRange1(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 左边(模板1)
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		//
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] != target {
		return []int{-1, -1}
	}
	start := r
	// 右边(模板2)
	l, r := 0, len(nums)-1
	for l < r {
		mid := (r + l + 1) >> 1
		//
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	end := r
	return []int{start, end}
}

/*
	二分法套路2
*/
func searchRange(nums []int, target int) []int {
	l := leftBound(nums, target)
	r := rightBound(nums, target)
	return []int{l, r}
}

func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// right:=len(nums)
	left, right := 0, len(nums)
	//
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] >= target {
			right = mid
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}
func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// right:=len(nums)
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] <= target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	if left-1 < 0 || nums[left-1] != target {
		return -1
	}
	return left - 1
}