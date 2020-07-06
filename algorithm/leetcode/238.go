/*
	给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

	你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

	思路1：动态规划（两个数组，一个存从左边开始到某个数的乘积，一个存从右边开始到某个数的乘积）
	思路2：当前解法
*/
func productExceptSelf(nums []int) []int {
	result:=make([]int,len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = 1
	}
	product:=1
	for i := 0; i < len(nums); i++ {
		// 到某个数为止，前面所有数的乘积
		result[i] = result[i]*product
		product = product*nums[i]
	}
	product=1
	for i := len(nums)-1; i>=0; i-- {
		// 到某个数为止，后面所有数的乘积
		result[i] = result[i]*product
		product = product*nums[i]
	}
	return result
}