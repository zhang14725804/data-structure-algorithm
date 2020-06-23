/*
	搜索旋转排序数组

	当nums元素个数较少的时候，直接for循环暴力枚举
	（1）先二分出两段的分界点
	（2）在判断target在那一段，，然后二分结果
*/
func search(nums []int, target int) int {
    if len(nums) == 0{
		return -1
	}
	// 👉：找到最小元素
	l,r := 0,len(nums)-1
	// 👉：这种通常题目描述为满足某种情况的最小的元素
	for l<r{
		mid:=(l+r) >> 1
		if nums[mid] <= nums[len(nums)-1]{
			r = mid
		}else{
			l = mid + 1
		}
	}

	if target <= nums[len(nums)-1]{
		// 在右边
		r = len(nums)-1
	}else{
		// 在左边
		l = 0
		r--
	}

	// 是否存在一个目标值（精确查找）
	for l<r{
		mid := (l+r) >> 1
		if nums[mid] >= target{
			r = mid
		}else{
			l = mid + 1
		}
	}

	if nums[l] == target{
		return l
	}
	return -1
}

func search(nums []int, target int) int {
    if len(nums) == 0{
		return -1
	}
	//  👉：找到最大元素
	l,r := 0,len(nums)-1
	// 👉：这种通常题目描述为满足某种情况的最大的元素
	for l<r{
		mid:=(l+r+1) >> 1
		if nums[mid] >= nums[0]{
			l = mid
		}else{
			r = mid - 1
		}
	}

	if target >= nums[0]{
		// 在左边
		l = 0
		r = l
	}else{
		// 在右边
		l = l+1
		r = len(nums)-1
	}

	// 是否存在一个目标值（精确查找）
	for l<r{
		mid := (l+r) >> 1
		if nums[mid] >= target{
			r = mid
		}else{
			l = mid + 1
		}
	}

	if nums[l] == target{
		return l
	}
	return -1
}