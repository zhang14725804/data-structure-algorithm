/*
	给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

	我的思路:for循环放到map里面,数字作为key,个数为value
	实现思路:抽屉原理
*/
func findDuplicate(nums []int) int {
	n := len(nums)-1
	// 
	l,r := 1,n
	for l < r{
		mid := (l+r) >> 1
		cnt := 0
		//
		for _,x := range nums{
			if x >= l && x <= mid{
				cnt++
			}
		}
		// 
		if cnt > mid-l+1{
			r = mid
		}else{
			l = mid + 1
		}
	}
	return r
}