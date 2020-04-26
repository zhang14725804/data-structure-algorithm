/*
	旋转数组的最小数字（没懂题意😅）
	todo：二分法
*/
func findMin(nums []int) int {
	n := len(nums)-1
	if n<0{
		return -1
	}
	// 去重
	for n>0 && nums[n]==nums[0]{
		n--
	}
	if nums[n]>=nums[0]{
		return nums[0]
	}
	// todo：二分算法，死循环
	l,r:=0,n
	for l<r {
		mid := l+r >> 1 // [1,mid],[mid+1,r]
		if nums[mid] < nums[0]{
			r = mid
		}else{
			l = mid+1
		}
	}
	return nums[r]
}
