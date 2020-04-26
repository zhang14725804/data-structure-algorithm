/*
	不修改数组找出重复数字，O(1)的额外空间
	todo：”分治+抽屉原理“思路，数组分为两半

	[2, 3, 5, 4, 3, 2, 6, 7]测试通过
	todo：[1, 7, 5, 9, 7, 9, 5, 1, 5, 6, 7]测试不通过
*/


func duplicateInArray(nums []int) int {
	// 数组长度n+1，（1~n）作为数组长度
	l,r := 1,len(nums)-1
	for l < r{
		mid:=l+r >> 1 // [1,mid],[mid+1,r]
		s := 0
		for _,x := range nums{
			if x>= l && x<=mid{
				s++
			}
		}
		if s > mid-l+1 {
			r = mid
		}else{
			l = mid+1
		}
	}
	return r
}