/*
	给你一个整数数组 arr 和一个整数 k ，其中数组长度是偶数，值为 n 。

	现在需要把数组恰好分成 n / 2 对，以使每对数字的和都能够被 k 整除。

	如果存在这样的分法，请返回 True ；否则，返回 False 。

	思路：只需要关心arr中的每个数对k取余，这些余数是从0~k-1,检查从1~k-1的余数是不是两两配对

	todo:有个10^9 这个没有处理，代码完全有问题
*/
func canArrange(nums []int, target int) bool {
	n := len(nums)
	p := make([]int, n)
	p[0] = 1

	sort(nums)
	res := 0
	for i, j := 0, n-1; i < n; i++ {
		for j >= i && nums[i]+nums[j] > target {
			j--
		}
		if j >= i && nums[i]+nums[j] <= target {
			res = (res + p[j-i])
		}
	}
	return res > 0
}