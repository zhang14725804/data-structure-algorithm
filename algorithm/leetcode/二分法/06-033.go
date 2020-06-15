/*
	 搜索旋转排序数组（不太好理解）
*/
func search(nums []int, target int) int {
	if len(nums) == 0{
		return -1
	}
	// 找到最小值
	l,r := 0,len(nums)-1
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

	for l<r{
		mid := (l+r) >> 1
		if nums[mid] >= target{
			r = mid
		}else{
			l = mid + 1
		}
	}

	// 注意:这里只能用l
	if nums[l] == target{
		return l
	}
	return -1
}