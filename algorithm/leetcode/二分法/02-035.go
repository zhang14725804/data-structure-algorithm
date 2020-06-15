/*
	找到大于等于target得最左边的数（模板1）
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || nums[len(nums)-1]<target{
		return len(nums)
	}
	l,r := 0,len(nums)-1
	for l<r {
		mid := (l+r) >> 1
		// 
		if nums[mid] >= target{
			r = mid
		}else{
			l = mid+1
		}
	}
	return l
}

// 😅
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || nums[len(nums)-1]<target{
		return len(nums)
	}
	l,r := 0,len(nums)-1
	for l<r {
		mid := (l+r) >> 1
		// 
		if nums[mid] < target{
			l = mid+1
		}else{
			r = mid
		}
	}
	return l
}