/*
	输入一个 非空 整型数组，数组里的数可能为正，也可能为负。

	数组中一个或连续的多个整数组成一个子数组。

	求所有子数组的和的最大值。

	要求时间复杂度为O(n)。

	todo：思路不错
*/
func maxSubArray(nums []int) int {
	res,sum:= -1000000,0
	// 遍历当前数据
	for _,x := range nums{
		// todo：这几个意思
		if sum < 0{
			sum = 0
		}
		sum += x
		res = max(res,sum)
	}
	return res
}
func max(x,y int) int{
	if x>y{
		return x
	}
	return y
}