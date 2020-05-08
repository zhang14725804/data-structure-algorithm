/*
	数组中数值和下标相等的元素
	思路：二分法（又是二分法😅）
*/
func getNumberSameAsIndex(nums []int) int {
	l,r := 0,len(nums)-1
	for l<r{
		// 要有括号
		mid := (l+r) >> 1
		if nums[mid]-mid >= 0 {
			r = mid
		}else{
			l = mid+1
		}
	}
	if nums[r] - r == 0 {
		return r
	}
	return -1
}
