/*
	给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	思路1：暴力破解
	思路2：动态规划
		（1）状态表示，包括集合（所有以i结尾的子段）、属性（Max）
		（2）状态计算 f[i] = max(f[i-1], 0) + nums[i]（f[i-1]取或者不取的最大值，再加上f[i]）
*/

func maxSubArray(nums []int) int {
	res := INT_MIN
	last := 0
	for i := 0; i < len(nums); i++ {
		//
		now := compare(last, 0, true) + nums[i]
		res = compare(res, now, true)
		last = now
	}
	return res
}

/*
	暴力枚举，数量太大就不行了😅
*/
func maxSubArray(nums []int) int {
	res := INT_MIN
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			res = compare(res, sum, true)
		}
	}
	return res
}