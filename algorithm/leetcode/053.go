package leetcode

import (
	"math"
)

// Leetcode053 Maximum Subarray 动态规划方法
func Leetcode053(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	/*
		float64 初始化res, last, zero
		{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		{-1, -5}
		{-1, 0, -2}
		{-2, 1, -3, 4, -1, 2, 1, -5, 4}

		f[i] = max(f[i-1],0) + nums[i]
	*/
	res, last, zero := float64(0), float64(0), float64(0)
	for i := 0; i < len(nums); i++ {
		now := math.Max(last, zero) + float64(nums[i])
		// 排除都是负数和0的情况
		if i == 0 && res == 0 && now < 0 {
			res = now
		} else {
			res = math.Max(res, now)
		}
		last = now
	}
	return int(res)
}
