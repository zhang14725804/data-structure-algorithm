/*
	0到n-1中缺失的数字
	一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0到n-1之内
	思路1：二分法
	思路2：高斯求和公式
*/
func getMissingNumber(nums []int) int {
	//
	n := len(nums)+1
	res := n*(n-1)/2
	for _,x := range nums{
		res -= x
	}
	return res
}