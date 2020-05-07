/*
	数字在排序数组中出现的次数
	思路1：hash+循环
	思路2：二分法，难😅。二分法有两种划分方式（取mid方式不同，容易造成死循环）
*/
func getNumberOfK(nums []int , k int) int{
	if len(nums) == 0{
		return 0
	}
	// 第一种二分法
	l,r := 0,len(nums)-1
	for l<r{
		// 
		mid := (l+r) >> 1
		if nums[mid] < k{
			l = mid+1
		} else{
			r =mid
		}
	}
	if nums[l] != k{
		return 0
	}
	left := l

	// 第二种二分法
	l,r = 0,len(nums)-1
	for l<r{
		mid := (l+r+1) >> 1
		if nums[mid] <= k{
			l = mid
		}else{
			r = mid-1
		}
	}
	return r-left+1
}
