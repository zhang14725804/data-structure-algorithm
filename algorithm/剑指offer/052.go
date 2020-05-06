/*
	数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字
	假设要求只能使用 O(n) 的时间和额外 O(1) 的空间，该怎么做呢？
	todo：思路清奇
*/
func moreThanHalfNum_Solution(nums []int) int {
	cnt,val := 0,-1
	for _,x := range nums{
		if cnt == 0 {
			val = x
			cnt = 1
		}else{
			if x == val{
				cnt++
			}else{
				cnt--
			}
		}
	}
	return val
}