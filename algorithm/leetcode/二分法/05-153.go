/*
	寻找旋转排序数组中的最小值.假设按照升序排序的数组在预先未知的某个点上进行了旋转。
	todo:不好理解
*/
func findMin(nums []int) int {
	l,r := 0,len(nums)-1
	for l<r {
		mid := (l+r) >> 1
		// 模板1
		if nums[mid] <= nums[len(nums)-1]{
			r = mid
		}else{
			l = mid+1
		}
	}
	return nums[l]
}

func findMin(nums []int) int {
	l,r := 0,len(nums)-1
	for l<r {
		mid := (l+r) >> 1
		// 模板1
		if nums[mid] >= nums[0]{
			l = mid+1
		}else{
			r = mid
		}
	}
	return nums[l]
}