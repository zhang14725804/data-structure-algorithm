/*
	假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖 一次 该股票可能获得的利润是多少？

	枚举（枚举在那天卖）
*/
func maxDiff(nums []int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}
	// 记录最大值和最小值
	for i, minv := 1, nums[0]; i < len(nums); i++ {
		res = max(res, nums[i]-minv)
		minv = min(minv, nums[i])
	}
	return res
}